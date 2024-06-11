/* eslint-disable no-console */
const Generator = require('yeoman-generator');
const path = require('node:path');

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
        message: 'Enter the resource model triplet (in the form namespace:repo-name:model-name):',
      },
      {
        type: 'input',
        name: 'apiName',
        message: 'which API does the module implement?',
      },
    ];

    this.models = [];
    this.apis = [];

    async function doPrompts(resolve) {
      await this.prompt(prompts).then((props) => {
        const names = props.triplet.split(':');
        this.nameSpace = names?.[0];
        this.moduleName = names?.[1];
        this.models.push(names?.[2]);
        this.apis.push(props.apiName);
      });
      this.numModels -= 1;
      if (this.numModels > 0) {
        doPrompts.call(this, resolve);
      } else {
        resolve();
      }
    }
    new Promise((r) => {
      doPrompts.call(this, r);
    }).then(() => {
      cb();
    });
  }

  writing() {
    console.log('\nGenerating modules');

    let i = 0;

    // Get function stubs from the go SDK.
    while (i < this.apis.length) {
            const clientPath = path.join(__dirname, '/../viam-sdk/components', this.apis[i], '/client.go');
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
        functions[j] = `func (s *${this.moduleName}${this.models[i]})` + functions[j];

        // remove code inside of function
        const inside = functions[j].substring(functions[j].indexOf('{\n') + 1, functions[j]?.lastIndexOf('}'));
        functions[j] = functions[j].replace(inside, '\n\n');

        // Theres a bug where if the function returns a struct in the component package, it needs the package name to import it
        // Properties is a common one but there are others that one may need to add manually.
        const index =  functions[j].lastIndexOf('Properties');
        if (index != -1) {
          functions[j] = functions[j].slice(0, index) + this.apis[i] + '.' + functions[j].slice(index);
        }
        j+=1
      }

      // Get apiname with a capital letter to use in the template.
      let apiNameUppercase = this.apis[i][0].toUpperCase() + this.apis[i].slice(1);
      // powersensor and movementsensor need two letters capatilized.
      const index =  apiNameUppercase.indexOf('sensor');
      if (index != -1) {
          apiNameUppercase = apiNameUppercase.slice(0, index) + apiNameUppercase[index].toUpperCase() + apiNameUppercase.slice(index + 1);
      }

      const tmplContext = {
        moduleName: this.moduleName,
        nameSpace: this.nameSpace,
        apiName: this.apis[i],
        modelName: this.models[i],
        apiNameUppercase: apiNameUppercase,
        funcs: functions.join(''),
        moreImports: imports,
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
