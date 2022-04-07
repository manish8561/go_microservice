import web3Helper from "../../helpers/common/web3.helper";
import pairContractABI from '../../bin/pairContractABI.json';
import TokenABI from '../../bin/tokenContract.ABI.json';
import MasterchefABI from '../../bin/masterChefContractABI.json';
import StrategyPair from '../../bin/strategy.pairABI.json';
import StrategyToken from '../../bin/strategy.singleABI.json';
import axios from "axios";

class MasterChef extends web3Helper {
  constructor() {
    super();
  }

  public async calculateAPY(apr: any): Promise<string> {
    try {
      //  * @param interest {Number} APR as percentage (ie. 5.82)
      //  * @param frequency {Number} Compounding frequency (times a year)
      //  * @returns {Number} APY as percentage (ie. 6 for APR of 5.82%)
      const interest: any = apr
      // const SECONDS_PER_YEAR = 365.25 * 24 * 60 * 60;
      const BLOCKS_IN_A_YEAR = 28800
      const aprToApy: any = ((1 + (interest / 100)) ** (1 / BLOCKS_IN_A_YEAR) - 1) * BLOCKS_IN_A_YEAR * 100;
      return aprToApy
    } catch (err) {
      throw err;
    }
  }

  public async calculateTVLValue(deposit_token: string, strategyAddress: string, token_type: string): Promise<string> {
    try {
      let ABI: any;
      if (token_type === 'pair') {
        ABI = StrategyPair;
      } else if(token_type === 'native'){
        ABI = StrategyPair;
      } else {
        ABI = StrategyToken;
      }
      const contract: any = await this.callPairContract(ABI, strategyAddress);
      let tvl: any = await contract.methods.totalDeposits().call();
      const decimalVal: any = await contract.methods.decimals().call();
      const dollerPrice: any = await this.calPrice2(deposit_token);
      tvl = (tvl / 10 ** decimalVal) * Number(dollerPrice);
      return tvl.toFixed(2);
    } catch (err) {
      throw err;
    }
  }

  public async calculateAPRValue(masterChefAddress: string, lp: string): Promise<string> {
    try {
      let ACPrice = 0
      const ACToken: any = await axios.get(`${process.env.FARM_API_URL}pricefeeds?symbol=AC`);
      if (ACToken.status = 200) { ACPrice = ACToken.data.data.price }
      const totalAllcationPoint: any = await this.totalAllocationPoint(masterChefAddress);
      const allocationPoint: any = await this.allocationPoint(1, masterChefAddress);
      const acPerBlock: any = await this.acPerBlock(masterChefAddress);
      const liquidity: any = await this.handleLiquidity(lp, masterChefAddress)
      const apr: any = ((allocationPoint.allocPoint / totalAllcationPoint) * ((acPerBlock / 10 ** 18) * 28800 * 365 * 100 * ACPrice)) / liquidity;
      return apr.toFixed(4);
    } catch (err) {
      throw err;
    }
  }

  public async handleLiquidity(tokenAddress: any, contractAddress: any): Promise<Number> {
    try {
      if (tokenAddress != "0x0000000000000000000000000000000000000000") {
        const d: any = await this.getTokenDeposit(tokenAddress, contractAddress);
        let tokenPrice: any = await this.calPrice(tokenAddress)
        return d * tokenPrice
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
      const contract: any = await this.callPairContract(pairContractABI, pairAddress);
      return await contract.methods.token0().call();
    } catch (error) {
      throw error;
    }
  };

  public async getTokenOne(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callPairContract(pairContractABI, pairAddress);
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

  public async getDecimal(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(TokenABI, pairAddress);
      return await contract.methods.decimals().call();
    } catch (error) {
      throw error;
    }
  };

  public async calPrice(pairAddress: any): Promise<Number> {
    try {
      let price = 0
      let priceTokenZero: any = 0;
      let priceTokenOne: any = 0;
      let tokenZero: any;
      if (pairAddress === "") {
        return 0;
      }
      try {
        tokenZero = await this.getTokenZero(pairAddress);
      } catch (err) {
        // console.log('not a pair error', err);
        const symbolSingle = await this.getSymbol(pairAddress);
        const respTokenOne = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolSingle}`);
        if (respTokenOne.status === 200) {
          return respTokenOne.data.data.price
        }
        return 0;
      }

      if (tokenZero === '0x0000000000000000000000000000000000000000') {
        const symbolSingle = await this.getSymbol(pairAddress);
        const respTokenOne = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolSingle}`);
        if (respTokenOne.status === 200) {
          return respTokenOne.data.data.price
        }
        return 0;
      } else {
        const tokenOne: any = await this.getTokenOne(pairAddress);
        const reserve: any = await this.getReserves(pairAddress);
        const symbolZero: any = await this.getSymbol(tokenZero);
        const symbolOne: any = await this.getSymbol(tokenOne);
        const decimalZero: any = await this.getDecimal(tokenZero);
        const decimalOne: any = await this.getDecimal(tokenOne);
        // fetching data from Api for token zero...
        const respTokenZero = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolZero}`);
        if (respTokenZero.status === 200) {
          if (respTokenZero.data.data.symbol === symbolZero) {
            priceTokenZero = respTokenZero.data.data.price * reserve[0] / 10 ** decimalZero;
          }
        }
        // fetching data from Api for token one...
        const respTokenOne = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolOne}`);
        if (respTokenOne.status === 200) {
          if (respTokenOne.data.data.symbol === symbolOne) {
            priceTokenOne = respTokenOne.data.data.price * reserve[1] / 10 ** decimalOne
          }
        }
        price = priceTokenZero + priceTokenOne
        return price
      }
    } catch (err) {
      throw err;
    }
  }

  public async calPrice2(pairAddress: any): Promise<Number> {
    try {
      let price = 0
      let priceTokenZero: any = 0;
      let priceTokenOne: any = 0;
      let tokenZero: any;
      if (pairAddress === "") {
        return 0;
      }
      try {
        tokenZero = await this.getTokenZero(pairAddress);
      } catch (err) {
        // console.log('not a pair error', err);
        const symbolSingle = await this.getSymbol(pairAddress);
        const respTokenOne = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolSingle}`);
        if (respTokenOne.status === 200) {
          return respTokenOne.data.data.price
        }
        return 0;
      }

      if (tokenZero === '0x0000000000000000000000000000000000000000') {
        const symbolSingle = await this.getSymbol(pairAddress);
        const respTokenOne = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolSingle}`);
        if (respTokenOne.status === 200) {
          return respTokenOne.data.data.price
        }
        return 0;
      } else {
        const tokenOne: any = await this.getTokenOne(pairAddress);
        const symbolZero: any = await this.getSymbol(tokenZero);
        const symbolOne: any = await this.getSymbol(tokenOne);
        // fetching data from Api for token zero...
        const respTokenZero = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolZero}`);
        if (respTokenZero.status === 200) {
          if (respTokenZero.data.data.symbol === symbolZero) {
            priceTokenZero = respTokenZero.data.data.price;
          }
        }
        // fetching data from Api for token one...
        const respTokenOne = await axios(`${process.env.FARM_API_URL}pricefeeds?symbol=${symbolOne}`);
        if (respTokenOne.status === 200) {
          if (respTokenOne.data.data.symbol === symbolOne) {
            priceTokenOne = respTokenOne.data.data.price;
          }
        }
        price = priceTokenZero + priceTokenOne
        return price
      }
    } catch (err) {
      throw err;
    }
  }
}

export default new MasterChef();
