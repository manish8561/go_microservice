
//chainId supported
export const chainIdArr = [4];

export const blockDiff = 200;

//block mined
export const network: any = {
    1: { //eth
        blockMined: 5760,
    },
    4: {//rinkeby
        blockMined: 5760,
    },
    56: { //bsc
        blockMined: 28800,
    },
    97: {//bsc testnet
        blockMined: 28800,
    },
    137: {//polygon mainnet
        blockMined: 0.0004
    },
    80001: {//polygon testnet
        blockMined: 43200
    }
};
