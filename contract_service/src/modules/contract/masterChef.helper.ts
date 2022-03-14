import web3Helper from "../../helpers/common/web3.helper";
import pairContractABI from '../../bin/pairContractABI.json'
import TokenABI from '../../bin/tokenContract.ABI.json'
import MasterchefABI from '../../bin/masterChefContractABI.json'
import PairABI from '../../bin/pairContractABI.json'

// import tokenStrategyABI from '../../bin/strategy.singleABI.json'
// import pairStrategyABI from '../../bin/strategy.pairABI.json'

import axios from "axios";

class MasterChef extends web3Helper {
  constructor() {
    super();
  }
    public async calculateAPY(deposit_token: string, strategyAddress: string): Promise<string> {
      try {
        /**
        //  *
        //  * @param interest {Number} APR as percentage (ie. 5.82)
        //  * @param frequency {Number} Compounding frequency (times a year)
        //  * @returns {Number} APY as percentage (ie. 6 for APR of 5.82%)
        //  */
        const interest: any = await this.calculateAPRValue(deposit_token , strategyAddress)
        const SECONDS_PER_YEAR = 365.25 * 24 * 60 * 60;
        const frequency = SECONDS_PER_YEAR / 14;
        // const  frequency = /
        const aprToApy: any =  ((1 + (interest / 100)) ** (1 / frequency) - 1) * frequency * 100;
        return aprToApy
      } catch (err) {
        throw err;
      }
    }

  public async calculateTVLValue(deposit_token: string, strategyAddress: string): Promise<string> {
    try {
      const contract: any = await this.callContract(PairABI, deposit_token);
      const resp: any = await contract.methods.balanceOf(strategyAddress).call();
      console.log("resp", resp)
      return resp
    } catch (err) {
      throw err;
    }
  }

  public async calculateAPRValue(masterChefAddress: string, lp: string): Promise<string> {
    try {
      let ACPrice = 0

      console.log("ACTokenACTokenACTokenACToken", process.env.FARM_API_URL)


      const ACToken: any = await axios.get(`${process.env.FARM_API_URL}pricefeeds?symbol=AC`);

      if (ACToken.status = 200) { ACPrice = ACToken.data.data.price }
      const totalAllcationPoint: any = await this.totalAllocationPoint(masterChefAddress);
      const allocationPoint: any = await this.allocationPoint(1, masterChefAddress);
      const acPerBlock: any = await this.acPerBlock(masterChefAddress);
      const liquidity: any = await this.handleLiquidity(lp, masterChefAddress)
      const apr: any = ((allocationPoint.allocPoint / totalAllcationPoint) * ((acPerBlock / 10 ** 18) * 28800 * 365 * 100 * ACPrice)) / liquidity;
      return apr;
    } catch (err) {
      throw err;
    }
  }


  public async handleLiquidity(tokenAddress: any, contractAddress: any): Promise<Number> {
    try {
      if (tokenAddress != "0x0000000000000000000000000000000000000000") {
        const d: any = await this.getTokenDeposit(tokenAddress, contractAddress);
        const respTokenOne = await axios.get(`${process.env.FARM_API_URL}pricefeeds?symbol=USDT`);
        if (respTokenOne.status === 200) {
          return d * respTokenOne.data.data.price
        }
        else return 0
      }
      return 0
    } catch (error) {
      throw error;
    }
  }

  public async acPerBlock(contractAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(MasterchefABI, contractAddress);
      return await contract.methods.cakePerBlock().call();
    } catch (error) {
      throw error;
    }
  };

  public async getTokenDeposit(pairAddress: any, masterChefAddress: any): Promise<Number> {
    try {
      const contract: any = await this.callContract(TokenABI, pairAddress);
      const decimals = await contract.methods.decimals().call();
      let result = await contract.methods.balanceOf(masterChefAddress).call()
      result = (Number(result) / 10 ** decimals).toFixed(5);
      return Number(result);
    } catch (error) {
      throw error;
    }
  };


  public async allocationPoint(index: any, contractAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(MasterchefABI, contractAddress);
      const result = await contract.methods.poolInfo(index).call();
      return result
    } catch (error) {
      throw error;
    }
  };

  public async totalAllocationPoint(contractAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(MasterchefABI, contractAddress);
      return await contract.methods.totalAllocPoint().call();
    } catch (error) {
      throw error;
    }
  };

  public async getTokenZero(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(pairContractABI, pairAddress);
      const resp = await contract.methods.token0().call();
      return resp
    } catch (error) {
      return '0';
    }
  };

  public async getTokenOne(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(pairContractABI, pairAddress);
      return await contract.methods.token1().call();
    } catch (error) {
      throw error;
    }
  };

  public async getReserves(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(pairContractABI, pairAddress);
      return await contract.methods.getReserves().call();
    } catch (error) {
      throw error;
    }
  };

  public async getDecimals(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(TokenABI, pairAddress);
      return await contract.methods.decimals().call();
    } catch (error) {
      throw error;
    }
  };

  public async getSymbol(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(TokenABI, pairAddress);
      return await contract.methods.symbol().call();
    } catch (error) {
      throw error;
    }
  };

  public async calPrice(pairAddress: any): Promise<Number> {
    try {
      let price = 0
      let priceTokenZero: any = 0;
      let priceTokenOne: any = 0;
      if (pairAddress === "0x0000000000000000000000000000000000000000") {
        return 0;
      }
      const tokenZero: any = await this.getTokenZero(pairAddress);

      if (tokenZero === '0') {
        const symbolSingle = await this.getSymbol(pairAddress);
        const respTokenOne = await axios(`${process.env.FARM_API_URL}${symbolSingle}`);
        if (respTokenOne.status === 200) {
          if (respTokenOne.data.data.symbol === symbolSingle) {
            return respTokenOne.data.data.price
          }
        }
        return 0
      }
      else {
        const tokenOne: any = await this.getTokenOne(pairAddress);
        const reserve: any = await this.getReserves(pairAddress);
        const symbolZero = await this.getSymbol(tokenZero);
        const symbolOne = await this.getSymbol(tokenOne);

        // fetching data from Api for token zero...
        const respTokenZero = await axios(`${process.env.FARM_API_URL}${symbolZero}`);
        if (respTokenZero.status === 200) {
          if (respTokenZero.data.data.symbol === symbolZero) {
            priceTokenZero = respTokenZero.data.data.price * reserve[0];
          }
        }

        // fetching data from Api for token one...
        const respTokenOne = await axios(`${process.env.FARM_API_URL}${symbolOne}`);
        if (respTokenOne.status === 200) {
          if (respTokenOne.data.data.symbol === symbolOne) {
            priceTokenOne = respTokenOne.data.data.price * reserve[1]
          }
        }

        price = priceTokenZero + priceTokenOne
        return price
      }
    } catch (err) {
      throw err
    }
  }


}

export default new MasterChef();
