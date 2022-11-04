const fs = require("fs");
const readline = require('readline');
const nunjucks = require("nunjucks");


async function processHoldersConf() {
    const fileStream = fs.createReadStream(__dirname + '/init_holders.conf');

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    let Holders = [];
    for await (const line of rl) {
        // Each line in input.txt will be successively available here as `line`.
        let vs = line.split(",")
        if (vs[0] != 0) {
            Holders.push({
                address: vs[0],
            })
        }
    }
    return Holders
}

processHoldersConf().then(function(Holders) {
    const data = {
        Holders: Holders
    };
    const templateString = fs.readFileSync(__dirname + '/template/Holders.template').toString();
    const resultString = nunjucks.renderString(templateString, data);
    fs.writeFileSync(__dirname + '/init_holders.js', resultString);
})