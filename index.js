const pify = require('pify')
const fs = require('fs')
const meow = require('meow')
const resizeImg = require('resize-img')
const {readFile, writeFile, stat, mkdir} = pify(fs)
const path = require('path')

const sizes = [32, 57, 76, 96, 120, 128, 144, 152, 180, 195, 196, 228, 270, 558]

const cli = meow(`
  Usage
    $ favicon-generator <input>

  Options
    --output, -o Folder to output the favicons, by default it will be 'favicons'

  Examples
    $ favicon-generator logo.png
    $ favicon-generator logo.png --output images
`, {
  alias: {
    o: 'output'
  }
})

async function main(input, output = 'favicons') {
  const file = await readFile(path.join(process.cwd(), input))
  const outputDir = path.join(process.cwd(), output)

  try {
    await stat(outputDir)
  } catch(err) {
    await mkdir(outputDir)
  }

  for (let size of sizes) {
    const favicon = await resizeImg(file, {width: size, height: size})

    try {
      await writeFile(
        path.join(outputDir, `favicon-${size}.png`), favicon
      )
    } catch(err) {
      console.error(err.message)
    }
  }
}

main(cli.input[0], cli.flags.output)
