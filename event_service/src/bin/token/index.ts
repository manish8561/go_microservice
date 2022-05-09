
//chainId supported
export const chainIdArr = [4];

export const blockDiff = 200;

//events to be 
export const eventsArr = [
    "ProposalCreated",
    "StartBlockSet",
    "VoteCast",
    "ProposalCanceled",
    "ProposalQueued",
    "ProposalExecuted",
];
// intial values for events in the collections
export const initial: any = {
    "ProposalCreated": [
        { chainId: 4, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for rinkeby
        { chainId: 1, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for mainnet
    ],
    "StartBlockSet": [
        { chainId: 4, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for rinkeby
    ],
    "VoteCast": [
        { chainId: 4, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for rinkeby
    ],
    "ProposalCanceled": [
        { chainId: 4, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for rinkeby
    ],
    "ProposalQueued": [
        { chainId: 4, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for rinkeby
    ],
    "ProposalExecuted": [
        { chainId: 4, transactionHash: "", blockNumber: 10645509, lastBlockNumber: 10645509, contractName: "governance", contract: "0x551a6eA97Fcd0cb9AB5E806D330dd5Db47eFAe74" },// for rinkeby
    ],
}