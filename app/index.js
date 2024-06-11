/* eslint-disable no-console */
const Generator = require('yeoman-generator');
const path = require('node:path');

const { spawnSync } = require('child_process');
const { exit } = require('process');

module.exports = class extends Generator {
  async prompting() {
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

    await this.prompt(numModel).then((props) => {
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
      console.log(this.apis);
      cb();
    });
  }

  writing() {
    console.log('\nGenerating modules');

    let j = 0;

    // get stuff from the viam sdk
    while (j < this.apis.length) {
      const clientPath = path.join(__dirname, '/../viam-sdk/components', this.apis[j], '/client.go');
      const clientCode = this.fs.read(clientPath);

      let imports = clientCode.substring(clientCode.indexOf('(\n') + 1, clientCode.indexOf(')'));

      // get rid of the unneeded imports
      imports = imports.replace('"go.viam.com/rdk/protoutils"\n', '');
      imports = imports.replace('commonpb "go.viam.com/api/common/v1"\n', '');
      imports = imports.replace(`pb "go.viam.com/api/component/${this.apis[j]}/v1"\n`, '');
      imports = imports.replace('"google.golang.org/protobuf/types/known/structpb"\n', '');
      imports = imports.replace('"go.viam.com/utils/rpc"\n', '');

      // get only the API functions
      const functions = clientCode.split('func').slice(2);
      let i = 0;
      while (i < functions.length) {
        // replace client with struct name
        functions[i] = functions[i].replace('(c *client)', `func (s *${this.moduleName}${this.models[j]})`);

        // remove code inside of function
        const inside = functions[i].substring(functions[i].indexOf('{\n') + 1, functions[i]?.lastIndexOf('}'));
        functions[i] = functions[i].replace(inside, '\n\n');
        i += 1;
      }

      const tmplContext = {
        moduleName: this.moduleName,
        nameSpace: this.nameSpace,
        apiName: this.apis[j],
        modelName: this.models[j],
        apiNameUppercase: this.apis[j].charAt(0).toUpperCase() + this.apis[j].slice(1),
        funcs: functions.join(''),
        moreImports: imports,
      };

      this.fs.copyTpl(
        this.templatePath('_module.go'),
        this.destinationPath(`${this.models[j]}/${this.models[j]}.go`),
        tmplContext,
      );

      j += 1;
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
