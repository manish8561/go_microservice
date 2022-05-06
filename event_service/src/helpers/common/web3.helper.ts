import Web3 from "web3";
import { blockDiff } from "../../bin/token";

class Web3Helper {
    public web3Obj: any;
    public contractAddress: any;
    public contractObj: any;
    public chainId: Number = 4;//rinkeby

    constructor() { }
    /**
     * select rpc according chainId
     * @param  {Number} chainId
     * @returns Promise
     */
    private async seletRPC(chainId: Number): Promise<string> {
        let rpc = 'https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161';

        switch (chainId) {
            case 56: //bsc mainnet
                rpc = process.env.RPC_BNB_URL!;
                break;
            case 4:
                rpc = process.env.RPC_RINKEBY_URL!;
                break;
            case 1: //eth mainnet
                rpc = process.env.RPC_MAINNET_URL!;
                break;
            default:
                rpc = process.env.RPC_RINKEBY_URL!;
        }
        return rpc;
    }

    /* dynamically call web3 */
    /**
     * @param  {number} chainId
     * @returns Promise
     */
    public async callWeb3(chainId: Number): Promise<any> {
        try {
            if (this.web3Obj && this.chainId === chainId) {
                return this.web3Obj;
            } else {
                this.chainId = chainId;
                const rpc = await this.seletRPC(chainId);
                this.web3Obj = new Web3(rpc)
                return this.web3Obj
            }
        } catch (error) {
            throw error;
        }
    };
    /**
     * get block number from network
     * @param  {number} chainId
     * @returns Promise
     */
    public async getBlockNumber(chainId: Number): Promise<number> {
        const web3Obj = await this.callWeb3(chainId);
        return await web3Obj.eth.getBlockNumber()
    }
    /* dynamically call the contract instance */
    /**
     * @param  {number} chainId
     * @param  {any} contractAbi
     * @param  {string} contractAddress
     * @returns Promise
     */
    public async callContract(chainId: number, contractAbi: any, contractAddress: string): Promise<string> {
        try {
            const web3Obj = await this.callWeb3(chainId);
            if (this.contractObj && this.contractAddress === contractAddress.toLowerCase()) {
                return this.contractObj;
            }
            this.contractAddress = contractAddress.toLowerCase();
            this.contractObj = new web3Obj.eth.Contract(contractAbi, contractAddress);
            return this.contractObj;
        } catch (error) {
            throw error;
        }
    };
    /**
     * get events fromBlock toBlock with eventName
     * @param  {any} data
     * @returns Promise
     */
    public async getEvents(data: any): Promise<any> {
        try {
            const { contractObj, event, fromBlock, toBlock } = data;
            const eventData = await contractObj.getPastEvents(event, {
                // Using an array means OR: e.g. 20 or 23
                fromBlock,
                toBlock: (fromBlock + blockDiff)
            });
            // console.log("Event =========", eventData.length, event);
            if (eventData.length > 0) {
                return eventData;
            }
            return [];
        } catch (error) {
            throw error;
        }
    }

    /* dynamically call the pair contract instance */
    // public async callPairContract(contractAbi: any, contractAddress: string): Promise<string> {
    //     try {
    //         this.contractAddress = contractAddress.toLowerCase();
    //         this.contractObj = new this.web3Obj.eth.Contract(contractAbi, contractAddress);
    //         return this.contractObj
    //     } catch (error) {
    //         throw error;
    //     }
    // };

}

export default new Web3Helper();
