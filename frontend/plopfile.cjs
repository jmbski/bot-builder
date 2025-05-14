/* module.exports = function (plop) {
    plop.setGenerator('component', {
      description: 'Generate a Vue component',
      prompts: [{
        type: 'input',
        name: 'name',
        message: 'Component name?',
      }],
      actions: [{
        type: 'add',
        path: './{{pascalCase name}}.vue',
        templateFile: 'plop-templates/component.hbs',
      }],
    });
  };
   */
const path = require('path')
const fs = require('fs')
/**
 * Lists files and directories in the given path.
 * @param {string} targetPath - The path to list the contents of.
 * @returns {string[]} - An array of file and directory names.
 */
function listDir(targetPath) {
	try {
		return fs.readdirSync(targetPath)
	} catch (err) {
		console.error(`Error reading directory: ${err.message}`)
		return []
	}
}

function getGenerator(plop, name) {
	return plop.setGenerator(name, {
		description: 'Generate a Vue component in the current directory',
		prompts: [
			{
				type: 'input',
				name: 'name',
				message: 'Component name?',
			},
		],
		actions: (data) => {
			const cwd = process.cwd()
			const filename = plop.getHelper('pascalCase')(data.name) + '.vue'
			const fullPath = path.join(cwd, filename)

			return [
				{
					type: 'add',
					path: fullPath,
					templateFile: `plop-templates/${name}.hbs`,
				},
			]
		},
	})
}
module.exports = function (plop) {
	/* plop.setGenerator('component', {
		description: 'Generate a Vue component in the current directory',
		prompts: [
			{
				type: 'input',
				name: 'name',
				message: 'Component name?',
			},
		],
		actions: (data) => {
			const cwd = process.cwd()
			const filename = plop.getHelper('pascalCase')(data.name) + '.vue'
			const fullPath = path.join(cwd, filename)

			return [
				{
					type: 'add',
					path: fullPath,
					templateFile: 'plop-templates/component.hbs',
				},
			]
		},
	}) */
    listDir(path.join(__dirname, 'plop-templates')).forEach(filename => {
        const name = filename.split('.')[0];
        getGenerator(plop, name);
    })
}