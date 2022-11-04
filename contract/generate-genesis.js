const { spawn } = require("child_process")
const program = require("commander")
const nunjucks = require("nunjucks")
const fs = require("fs")
const web3 = require("web3")

const validators = require("./validators")
const init_holders = require("./init_holders")


program.option("--bscChainId <bscChainId>",
    "bscChainId",
    "378");
program.option("-c, --chainid <chainid>", "chain id", "888")

program.option(
    "--initValidatorSetBytes <initValidatorSetBytes>",
    "initValidatorSetBytes",
    ""
)

program.option(
    "--initBurnRatio <initBurnRatio>",
    "initBurnRatio",
    "0"
)

program.option("--initConsensusStateBytes <initConsensusStateBytes>",
    "init consensusState bytes, hex encoding, no prefix with 0x",
    "42696e616e63652d436861696e2d4e696c650000000000000000000000000000000000000000000229eca254b3859bffefaf85f4c95da9fbd26527766b784272789c30ec56b380b6eb96442aaab207bc59978ba3dd477690f5c5872334fc39e627723daa97e441e88ba4515150ec3182bc82593df36f8abb25a619187fcfab7e552b94e64ed2deed000000e8d4a51000");


require("./generate-system");
require("./generate-systemReward");
require("./generate-validatorset");
require("./generate-tokenhub");
require("./generate-tendermintlightclient");
require("./generate-relayerincentivizecontract");
require("./generate-crosschain");
require("./generate-slash");

program.version("0.0.1")
program.option(
    "-o, --output <output-file>",
    "Genesis json file",
    "./genesis.json"
)
program.option(
    "-t, --template <template>",
    "Genesis template json",
    "./template/genesis.template"
)
program.parse(process.argv)

//compile contract

// function compileContract(key, contractFile, contractName) {
//     return new Promise((resolve, reject) => {
//         const ls = spawn("solc", [
//             "--bin",
//             "--abi",
//             // "--bin-runtime",
//             // "/=/",
//             // "--optimize",
//             // "--optimize-runs",
//             // "200",
//             contractFile,
//             "-o",
//             "./out"
//         ])

//         // const ls = spawn("solc", [
//         //     "--bin-runtime",
//         //     "/=/",
//         //     "--optimize",
//         //     "--optimize-runs",
//         //     "200",
//         //     contractFile
//         // ])

//         const result = []
//         ls.stdout.on("data", data => {
//             result.push(data.toString())
//         })

//         ls.stderr.on("data", data => {
//             // console.log(`stderr: ${data}`)
//         })

//         ls.on("close", code => {
//             // console.log(`child process exited with code ${code}`)
//             resolve(result.join(""))
//         })
//     }).then(compiledData => {
//         // compiledData = compiledData.replace(
//         //     `======= ${contractFile}:${contractName} =======\nBinary of the runtime part:`,
//         //     "@@@@"
//         // )
//         // const matched = compiledData.match(/@@@@\n([a-f0-9]+)/)
//         // return { key, compiledData: matched[1], contractName, contractFile }
//         const templateString = fs.readFileSync(`./out/${contractName}.bin`).toString()
//         return { key, compiledData: templateString, contractName, contractFile }
//     })
// }

function compileContract(key, contractFile, contractName) {
    return new Promise((resolve, reject) => {
        const ls = spawn("solc", [
            "--bin-runtime",
            "/=/",
            "--optimize",
            "--optimize-runs",
            "200",
            contractFile
        ])

        const result = []
        ls.stdout.on("data", data => {
            result.push(data.toString())
        })

        ls.stderr.on("data", data => {
            // console.log(`stderr: ${data}`)
        })

        ls.on("close", code => {
            console.log(`child process exited with code ${code}`)
            resolve(result.join(""))
        })
    }).then(compiledData => {
        console.log(compiledData)
        compiledData = compiledData.replace(
            `======= ${contractFile}:${contractName} =======\nBinary of the runtime part:`,
            "@@@@"
          )
          const matched = compiledData.match(/@@@@\n([a-f0-9]+)/)
        return { key, compiledData: matched[1], contractName, contractFile }
    })
}
// compile files
Promise.all([
    compileContract(
        "validatorContract",
        "sol/BSCValidatorSet.sol",
        "BSCValidatorSet"
    ),
    // compileContract(
    //     "systemRewardContract",
    //     "sol/SystemReward.sol",
    //     "SystemReward"
    // ),
    // compileContract(
    //     "slashContract",
    //     "sol/SlashIndicator.sol",
    //     "SlashIndicator"
    // ),
    // compileContract(
    //     "tendermintLightClient",
    //     "sol/TendermintLightClient.sol",
    //     "TendermintLightClient"
    // ),
    // compileContract(
    //     "tokenHub",
    //     "sol/TokenHub.sol",
    //     "TokenHub"
    // ),
    // compileContract(
    //     "relayerHub",
    //     "sol/RelayerHub.sol",
    //     "RelayerHub"
    // ),
    // compileContract(
    //     "relayerIncentivize",
    //     "sol/RelayerIncentivize.sol",
    //     "RelayerIncentivize"
    // ),
    // compileContract(
    //     "govHub",
    //     "sol/GovHub.sol",
    //     "GovHub"
    // ),
    // compileContract(
    //     "tokenManager",
    //     "sol/TokenManager.sol",
    //     "TokenManager"
    // ),
    // compileContract(
    //     "crossChain",
    //     "sol/CrossChain.sol",
    //     "CrossChain"
    // )
]).then(result => {

    program.option("--initLockedBNBOnTokenHub <initLockedBNBOnTokenHub>",
        "initLockedBNBOnTokenHub",
        "176405560900000000000000000");
    const data = {
        initLockedBNBOnTokenHub: program.initLockedBNBOnTokenHub,
        chainId: program.chainid,
        initHolders: init_holders,
        extraData: web3.utils.bytesToHex(validators.extraValidatorBytes)
    }
    result.forEach(r => {
        console.log(r.key)
        data[r.key] = r.compiledData
    })
    const templateString = fs.readFileSync(program.template).toString()
    const resultString = nunjucks.renderString(templateString, data)
    fs.writeFileSync(program.output, resultString)

})