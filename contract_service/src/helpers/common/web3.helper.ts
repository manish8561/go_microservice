import Web3 from "web3";

class Web3Helper {
    public web3Obj: any;
    public rpcurl = 'https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161';
    public contractAddress: any;
    public contractObj: any;
    constructor() {
        this.callWeb3(this.rpcurl);
    }
    /* dynamically call web3 */
    public async callWeb3(RpcURL: string): Promise<string> {
        try {
            if (this.web3Obj && this.rpcurl === RpcURL) {
                return this.web3Obj
            } else {
                this.rpcurl = RpcURL;
                this.web3Obj = new Web3(RpcURL)
                return this.web3Obj
            }
        } catch (error) {
            throw error;
        }
    };
    /* dynamically call the contract instance */
    public async callContract(contractAbi: any, contractAddress: string, ): Promise<string> {
        try {
            if (this.contractObj && this.contractAddress === contractAddress.toLowerCase()) {
                return this.contractObj
            }
            this.contractAddress = contractAddress.toLowerCase();
            this.contractObj = new this.web3Obj.eth.Contract(contractAbi, contractAddress);
            return this.contractObj
        } catch (error) {
            throw error;
        }
    };
}

export default Web3Helper;
