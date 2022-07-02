import pairContractABI from '../../bin/pairContractABI.json';
import TokenABI from '../../bin/tokenContract.ABI.json';
import MasterchefABI from '../../bin/masterChefContractABI.json';
import StrategyPair from '../../bin/strategy.pairABI.json';
import StrategyToken from '../../bin/strategy.singleABI.json';
import axios from "axios";
import * as Helpers from '../../helpers';
import { network } from '../../bin/token';


class MasterChef {

  /**
   * calculate APY
   * @param  {any} apr
   * @param  {number} chainId
   * @returns Promise
   */
  public async calculateAPY(apr: any): Promise<string> {
    try {
      //  * @param interest {Number} APR as percentage (ie. 5.82)
      //  * @param frequency {Number} Compounding frequency (times a year)
      //  * @returns {Number} APY as percentage (ie. 6 for APR of 5.82%)
      const interest: any = apr;
      // const SECONDS_PER_YEAR = 365.25 * 24 * 60 * 60;
      const DAYS_IN_YEAR = 365;
      // const aprToApy: any = ((1 + (interest / 100)) ** (1 / DAYS_IN_YEAR) - 1) * DAYS_IN_YEAR * 100;
      const aprToApy = (DAYS_IN_YEAR * apr).toString();
      return aprToApy;
    } catch (err) {
      throw err;
    }
  }
  /**
   * calculate total value locked
   * @param  {string} deposit_token
   * @param  {string} strategyAddress
   * @param  {string} token_type
   * @param  {number} chainId
   * @returns Promise
   */
  public async calculateTVLValue(deposit_token: string, strategyAddress: string, token_type: string, chainId: number): Promise<string> {
    try {
      let ABI: any;
      if (token_type === 'pair' || token_type === 'stable_pair' || token_type === 'native') {
        ABI = StrategyPair;
      } else {
        ABI = StrategyToken;
      }
      const contract: any = await Helpers.Web3Helper.callContract(chainId, ABI, strategyAddress);
      let tvl: any = await contract.methods.totalDeposits().call();
      const decimalVal: any = await contract.methods.decimals().call();
      const dollerPrice: any = await this.calPrice2(deposit_token, chainId);
      tvl = (tvl / 10 ** decimalVal) * Number(dollerPrice);
      return tvl.toFixed(6);
    } catch (err) {
      throw err;
    }
  }
  /**
   * get price from other service
   * @param  {string} token
   * @returns Promise
   */
  public async getTokenPriceUSD(token: string): Promise<number> {
    try {
      const resp: any = await axios.get(`${process.env.FARM_API_URL}api/farm_service/pricefeeds?symbol=${token}`);
      if (resp.status = 200) {
        return resp.data.data.price;
      }
      return 0;
    } catch (error) {
      // console.log(error, 'get price error');
      return 0;
    }
  }
  /**
   * calculate APR value
   * @param  {string} masterChefAddress
   * @param  {string} lp
   * @param  {number} chainId
   * @returns Promise
   */
  public async calculateAPRValue(masterChefAddress: string, lp: string, chainId: number, pid: number): Promise<string> {
    try {
      //reward token price (cake for pancake)
      let acPrice = 0;
      acPrice = await this.getTokenPriceUSD('AC');

      const masterchefContract: any = await Helpers.Web3Helper.callContract(chainId, MasterchefABI, masterChefAddress);
      const totalAllcationPoint: any = await masterchefContract.methods.totalAllocPoint().call();

      const poolInfo: any = await await masterchefContract.methods.poolInfo(pid).call();

      const cakePerBlock: any = await masterchefContract.methods.cakePerBlock().call();

      const liquidity: any = await this.handleLiquidity(lp, masterChefAddress, chainId);
      // console.log({acPrice, totalAllcationPoint, cakePerBlock, liquidity});
      if (liquidity === 0) {
        return '0';
      }
      const blockMined = network[chainId].blockMined;
      //since it is in cake value
      const accCakePerShare = poolInfo.accCakePerShare / (10 ** 18);
      const apr: any = ((accCakePerShare / totalAllcationPoint) * ((cakePerBlock / 10 ** 18) * blockMined * 100 * acPrice)) / liquidity;
      return apr.toFixed(4);
    } catch (err) {
      throw err;
    }
  }
  /**
   * @param  {any} tokenAddress
   * @param  {any} contractAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async handleLiquidity(tokenAddress: any, contractAddress: any, chainId: number): Promise<Number> {
    try {
      if (tokenAddress != "0x0000000000000000000000000000000000000000") {
        const d: any = await this.getTokenDeposit(tokenAddress, contractAddress, chainId);
        let tokenPrice: any = await this.calPrice(tokenAddress, chainId);
        return d * tokenPrice;
      }
      return 0;
    } catch (error) {
      throw error;
    }
  }

  /**
   * Autocompound Per Block
   * @param  {string} actoken
   * @param  {string} strategy
   * @param  {number} chainId
   * @returns Promise
   */
  public async getTokenPerBlock(actoken: string, strategy: string, chainId: number): Promise<Number> {
    try {
      let contract: any = await Helpers.Web3Helper.callContract(chainId, StrategyPair, strategy);
      let tokenPerBlock = await contract.methods.autoCompoundTokenPerBlock().call();
      contract = await Helpers.Web3Helper.callContract(chainId, TokenABI, actoken);
      const decimals = await contract.methods.decimals().call();
      const result = Number(tokenPerBlock) / 10 ** decimals;
      return Number(result);
    } catch (error) {
      throw error;
    }
  };
  /**
   * get token deposit
   * @param  {any} pairAddress
   * @param  {any} masterChefAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getTokenDeposit(pairAddress: any, masterChefAddress: any, chainId: number): Promise<Number> {
    try {
      const contract: any = await Helpers.Web3Helper.callContract(chainId, TokenABI, pairAddress);
      const decimals = await contract.methods.decimals().call();
      let result = await contract.methods.balanceOf(masterChefAddress).call();
      result = (Number(result) / 10 ** decimals);
      return Number(result);
    } catch (error) {
      throw error;
    }
  };

  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getTokenZero(pairAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callPairContract(pairAddress, chainId);
      return await contract.methods.token0().call();
    } catch (error) {
      throw error;
    }
  };
  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getTokenOne(pairAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callPairContract(pairAddress, chainId);
      return await contract.methods.token1().call();
    } catch (error) {
      throw error;
    }
  };
  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getReserves(pairAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callContract(chainId, pairContractABI, pairAddress);
      return await contract.methods.getReserves().call();
    } catch (error) {
      throw error;
    }
  };
  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getDecimals(pairAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callContract(chainId, TokenABI, pairAddress);
      return await contract.methods.decimals().call();
    } catch (error) {
      throw error;
    }
  };
  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getSymbol(pairAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callContract(chainId, TokenABI, pairAddress);
      return await contract.methods.symbol().call();
    } catch (error) {
      throw error;
    }
  };
  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async getDecimal(pairAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callContract(chainId, TokenABI, pairAddress);
      return await contract.methods.decimals().call();
    } catch (error) {
      throw error;
    }
  };
  /**
   * call price decenteralize way
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async calPrice(pairAddress: any, chainId: number): Promise<Number> {
    try {
      let price = 0;
      let priceTokenZero: any = 0;
      let priceTokenOne: any = 0;
      let tokenZero: any;
      if (pairAddress === "") {
        return 0;
      }
      try {
        tokenZero = await this.getTokenZero(pairAddress, chainId);
      } catch (err) {
        // console.log('not a pair error', err);
        const symbolSingle = await this.getSymbol(pairAddress, chainId);
        return await this.getTokenPriceUSD(symbolSingle);
      }

      if (tokenZero === '0x0000000000000000000000000000000000000000') {
        const symbolSingle = await this.getSymbol(pairAddress, chainId);
        return await this.getTokenPriceUSD(symbolSingle);
      } else {
        const tokenOne: any = await this.getTokenOne(pairAddress, chainId);
        const reserve: any = await this.getReserves(pairAddress, chainId);
        const symbolZero: any = await this.getSymbol(tokenZero, chainId);
        const symbolOne: any = await this.getSymbol(tokenOne, chainId);
        const decimalZero: any = await this.getDecimal(tokenZero, chainId);
        const decimalOne: any = await this.getDecimal(tokenOne, chainId);
        // fetching data from Api for token zero...
        const respTokenZero = await this.getTokenPriceUSD(symbolZero);
        if (respTokenZero) {
          priceTokenZero = respTokenZero * reserve[0] / 10 ** decimalZero;
        }
        // fetching data from Api for token one...

        const respTokenOne = await this.getTokenPriceUSD(symbolOne);
        if (respTokenOne) {
          priceTokenOne = respTokenOne * reserve[1] / 10 ** decimalOne;
        }
        price = priceTokenZero + priceTokenOne;
        return price;
      }
    } catch (err) {
      throw err;
    }
  }
  /**
   * @param  {any} pairAddress
   * @param  {number} chainId
   * @returns Promise
   */
  public async calPrice2(pairAddress: any, chainId: number): Promise<Number> {
    try {
      let price = 0;
      let priceTokenZero: any = 0;
      let priceTokenOne: any = 0;
      let tokenZero: any;
      if (pairAddress === "") {
        return 0;
      }
      try {
        tokenZero = await this.getTokenZero(pairAddress, chainId);
      } catch (err) {
        // console.log('not a pair error', err);
        const symbolSingle = await this.getSymbol(pairAddress, chainId);
        return await this.getTokenPriceUSD(symbolSingle);
      }

      if (tokenZero === '0x0000000000000000000000000000000000000000') {
        const symbolSingle = await this.getSymbol(pairAddress, chainId);
        return await this.getTokenPriceUSD(symbolSingle);
      } else {
        const tokenOne: any = await this.getTokenOne(pairAddress, chainId);
        const symbolZero: any = await this.getSymbol(tokenZero, chainId);
        const symbolOne: any = await this.getSymbol(tokenOne, chainId);
        // fetching data from Api for token zero...
        const respTokenZero = await this.getTokenPriceUSD(symbolZero);
        if (respTokenZero) {
          priceTokenZero = respTokenZero;
        }
        // fetching data from Api for token one...
        const respTokenOne = await this.getTokenPriceUSD(symbolOne);
        if (respTokenOne) {
          priceTokenOne = respTokenOne;
        }
        price = priceTokenZero + priceTokenOne;
        return price;
      }
    } catch (err) {
      throw err;
    }
  }
}

export default new MasterChef();
