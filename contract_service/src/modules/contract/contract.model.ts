// import * as SqlString from "sqlstring";
import BaseModel from "../../model/base.model";
import PriceFeed from "../pricefeed/pricefeed.model";

import Web3 from "web3";

const factoryContract = require("../../bin/factoryContractABI.json");
const masterChefContract = require("../../bin/masterChefContractABI.json");
const pairContract = require("../../bin/pairContractABI.json");
const routerContract = require("../../bin/routerContractABI.json");
const tokenContract = require("../../bin/RSTokenABI.json");

class Contract extends BaseModel {
  public chainId = 1280;
  public web3Obj: any;
  public WAAAToken = "0x8AC9810F7e3d1AefC8d6d2eE35B0D3CB3EE476D7";
 
  public contractAddress = process.env.CONTRACT_ADDRESS;
  public defaultAddress = process.env.DEFAULT_ADDRESS!;
  public blockNumber = process.env.BLOCK_NUMBER!;
  public tokenAddress = process.env.TOKEN_CONTRACT_ADDRESS!;

  constructor() {
    super();
    this.callW3Objects(1280);
  }
  private async callW3Objects(chainId:number) {
    let rpc:any = process.env.MOONRABBIT_RPC!;
    
    if (!this.web3Obj || this.chainId !== chainId) {
      switch (chainId){
        case 1280:
        rpc = process.env.MOONRABBIT_RPC;
        break;
        default:
          rpc = process.env.MOONRABBIT_RPC
      }
      this.web3Obj = new Web3(rpc);
    }
    return this.web3Obj;
  }

  /**
   * get eth price
   * @returns Promise
   */
  public async getTotalPairs(chainId:number, factory:string): Promise<string> {
    try {
      const web3Obj = await this.callW3Objects(chainId);
      const factoryContractObj = new web3Obj.eth.Contract(factoryContract, factory);
     return await factoryContractObj.methods.allPairsLength().call();
    } catch (err) {
      console.log(err,'contract call')
      return '0';
    }
  }

  public async getTotalLiquidity(chainId:number, masterchef:string, pair:string): Promise<any> {
    try {
      let totalLiquidity;
      const web3Obj = await this.callW3Objects(chainId);
      // const masterchefObj = new web3Obj.eth.Contract(masterChefContract, masterchef);
      const pairObj = new web3Obj.eth.Contract(pairContract, pair);
      // const routerObj = new web3Obj.eth.Contract(routerContract, router);

      const lpBalance = await pairObj.methods.balanceOf(masterchef).call();
      const totalLPSupply = await pairObj.methods.totalSupply().call();

      const token0 = await pairObj.methods.token0().call();
      const token1 = await pairObj.methods.token1().call();

      const reserve = await pairObj.methods.getReserves().call();

      let rsToken;
      let rsReserves;

      if (token0 === this.WAAAToken) {//aaa
        rsToken = token1;
        rsReserves = reserve[0];
      } else if (token1 === this.WAAAToken) {//RS
        rsToken = token0;
        rsReserves = reserve[1];
      }

      

      if (rsToken) {
        const tokenObj = new web3Obj.eth.Contract(tokenContract, rsToken);
        const rsDecimalFactor = 10 ** (await tokenObj.methods.decimals().call());
        const tokenPrice = await PriceFeed.getPrice('moon-rabbit');

        totalLiquidity = 2 * ((rsReserves * lpBalance * tokenPrice['moon-rabbit'].usd) / (totalLPSupply * rsDecimalFactor));

        console.log("liquidity is: ", totalLiquidity, rsReserves, rsToken, tokenPrice, rsDecimalFactor);
        return totalLiquidity;
      } else {
        return '0';
      }

      

      
    } catch (err) {
      console.log(err, "Contract call");
      return '0';
    }
  }
}

export default new Contract();
