
//chainId supported
export const chainIdArr = [4];

export const blockDiff = 50;

//tokens according to chainId
// export const tokens = {
//     1: {
//         governance: {
//             address: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4",
//             abi: governanceAbi
//         },
//         staking: {
//             address: "0xc0c7eC10F5B957A221926F348218e46EF34a611d",
//             abi: governanceAbi
//         }
//     },
//     4: {
//         governance: {
//             address: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4",
//             abi: governanceAbi
//         },
//         staking: {
//             address: "0xc0c7eC10F5B957A221926F348218e46EF34a611d",
//             abi: governanceAbi
//         }
//     },
// };
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
        { chainId: 4, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for rinkeby
        { chainId: 1, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for mainnet
    ],
    "StartBlockSet": [
        { chainId: 4, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for rinkeby
    ],
    "VoteCast": [
        { chainId: 4, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for rinkeby
    ],
    "ProposalCanceled": [
        { chainId: 4, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for rinkeby
    ],
    "ProposalQueued": [
        { chainId: 4, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for rinkeby
    ],
    "ProposalExecuted": [
        { chainId: 4, transactionHash: "", blockNumber: 10605232, lastBlockNumber: 10605232, contractName: "governance", contract: "0x86fC708e761Ab76B761e5d64F3e91Ffe781EA4e4" },// for rinkeby
    ],
}