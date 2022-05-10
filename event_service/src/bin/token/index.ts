
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
        { chainId: 4, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for rinkeby
        { chainId: 1, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for mainnet
    ],
    "StartBlockSet": [
        { chainId: 4, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for rinkeby
    ],
    "VoteCast": [
        { chainId: 4, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for rinkeby
    ],
    "ProposalCanceled": [
        { chainId: 4, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for rinkeby
    ],
    "ProposalQueued": [
        { chainId: 4, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for rinkeby
    ],
    "ProposalExecuted": [
        { chainId: 4, transactionHash: "", blockNumber: 10651529, lastBlockNumber: 10651529, contractName: "governance", contract: "0x02F6A4bf2326ca24b164e790cf1D97a192fdb35A" },// for rinkeby
    ],
}