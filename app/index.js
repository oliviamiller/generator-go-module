/* eslint-disable no-console */
const Generator = require('yeoman-generator');
const path = require('node:path');
const fs = require('fs');

const { spawnSync } = require('child_process');
const { exit } = require('process');

module.exports = class extends Generator {
  prompting() {
    console.log(
      '\n'
        + '  *** viam module generator ***  \n'
        + '\n',
    );

    const cb = this.async();

    const numModel = [
      {
        type: 'input',
        name: 'numModels',
        message: 'How many models are in the module?',
      },
    ];

    this.prompt(numModel).then((props) => {
      this.numModels = props.numModels;
    });

    const prompts = [
      {
        type: 'input',
        name: 'triplet',
        message: 'Enter the resource model triplet (in the form namespace:family:model-name):',
      },
      {
        type: 'input',
        name: 'apiName',
        message: 'Which API does this model implement?',
      },
    ];

    this.models = [];
    this.apis = [];
    this.families = [];

    async function doPrompts(resolve, reject) {
      await this.prompt(prompts).then((props) => {
        if (props.apiName == 'board' ||  props.apiName == 'camera') {
          reject(props.apiName + ' is not supported yet');
        }
        else {
          const names = props.triplet.split(':');
          this.nameSpace = names?.[0];
          this.families.push(names?.[1]);
          this.models.push(names?.[2]);
          this.apis.push(props.apiName);
        }
      });
      this.numModels -= 1;
      if (this.numModels > 0) {
        doPrompts.call(this, resolve, reject);
      } else {
        resolve();
      }
    }
    new Promise((resolve, reject) => {
      doPrompts.call(this, resolve, reject);
    }).then(() => {
      cb();
    },
    (error) => {
      console.log(error);
    });
  }

  writing() {
    console.log('\nGenerating modules');
    let i = 0;

    // Get function stubs from the go SDK.
    while (i < this.apis.length) {
      let clientPath = '';
      this.resourceType = 'Component';
      clientPath = path.join(__dirname, '/../viam-sdk/components', this.apis[i], '/client.go');

      // if component doesn't exist check if its a service.
      if (!fs.existsSync(clientPath)) {
            clientPath = path.join(__dirname, '/../viam-sdk/services', this.apis[i], '/client.go');
            this.resourceType = 'Service';
        }
      const clientCode = this.fs.read(clientPath);

      let imports = clientCode.substring(clientCode.indexOf('(\n') + 1, clientCode.indexOf(')'));

      // get rid of the unneeded imports.
      // there may be others that are unneccessary for the module that will get imported.
      imports = imports.replace('"go.viam.com/rdk/protoutils"\n', '');
      imports = imports.replace('rprotoutils "go.viam.com/rdk/protoutils"\n', '')
      imports = imports.replace('rprotoutils', '')
      imports = imports.replace('"go.viam.com/utils/protoutils"\n', '')
      imports = imports.replace('commonpb "go.viam.com/api/common/v1"\n', '');
      imports = imports.replace(`pb "go.viam.com/api/component/${this.apis[i]}/v1"\n`, '');
      imports = imports.replace('"google.golang.org/protobuf/types/known/structpb"\n', '');
      imports = imports.replace('"go.viam.com/utils/rpc"\n', '');

      // get only the client functions
      const functions = clientCode.split('func (c *client)').slice(1);
      let j = 0;

      while (j < functions.length) {
        // replace client with struct name
        functions[j] = `func (s *${this.families[i]}${this.models[i]})` + functions[j];

        // remove code inside of function
        const inside = functions[j].substring(functions[j].indexOf('{\n') + 1, functions[j]?.lastIndexOf('}'));
        functions[j] = functions[j].replace(inside, '\n\n');

        // Theres a bug where if the function returns a struct in the component package, it needs the package name to import it
        // Properties is a common one but there are others that one may need to add manually.
        let index =  functions[j].lastIndexOf('Properties');
        if (index != -1) {
          functions[j] = functions[j].slice(0, index) + this.apis[i] + '.' + functions[j].slice(index);
        }
        index =  functions[j].lastIndexOf('Accuracy');
        if (index != -1) {
          functions[j] = functions[j].slice(0, index) + this.apis[i] + '.' + functions[j].slice(index);
        }
        j+=1
      }

      let apiNameUppercase = getApiNameUpperCase(this.apis[i])

      let objName = ''
      if (this.resourceType == 'Service') {
        objName = 'Service'
      } else {
         objName = apiNameUppercase
      }

      const tmplContext = {
        moduleName: this.families[i],
        nameSpace: this.nameSpace,
        apiName: this.apis[i],
        modelName: this.models[i],
        apiNameUppercase: apiNameUppercase,
        funcs: functions.join(''),
        moreImports: imports,
        objName: objName,
        resourceType: this.resourceType,
        resourceTypeLower: this.resourceType.toLowerCase(),
      };

      this.fs.copyTpl(
        this.templatePath('_module.go'),
        this.destinationPath(`${this.models[i]}/${this.models[i]}.go`),
        tmplContext,
      );

      i += 1;
    }

    const mainContext = {
      moduleName: this.moduleName,
      nameSpace: this.nameSpace,
      apis: this.apis,
      models: this.models,
      resourceTypeLower: this.resourceType.toLowerCase(),
    };

    this.fs.copyTpl(
      this.templatePath('_main.go'),
      this.destinationPath('main.go'),
      mainContext,
    );
  }

  install() {
    const child = spawnSync('go', ['mod', 'init', this.moduleName], {
      cwd: process.cwd(),
      env: process.env,
      stdio: [process.stdin, process.stdout, process.stderr],
      encoding: 'utf-8',
    });
    if (child.error) {
      console.log(`Cannot run go mod init command: ${child.stdout}`);
      exit(1);
    }
  }

  static end() {
    console.log(
      '\n'
        + '  viam module template created!  \n'
        + '\n',
    );
  }
};

function capitalize(name, index) {
  name = name.slice(0, index) +
  name[index].toUpperCase() +
  name.slice(index + 1);
  return name;
}

function getApiNameUpperCase(name) {
   // Get apiname with a capital letter to use in the template.
   name = name[0].toUpperCase() + name.slice(1);

   // posetracker, powersensor and movementsensor need two letters capatilized.
   index = name.indexOf('sensor');
   if (index != -1) {
       name = capitalize(name, index);
   }
   index = name.indexOf('tracker');
   if (index != -1) {
      name = capitalize(name, index);
   }
   return name
}

