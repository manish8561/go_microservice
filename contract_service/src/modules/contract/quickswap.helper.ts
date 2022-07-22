import StakingRewardsABI from '../../bin/polygon/stakingRewards.json';
import StakingDualRewardsABI from '../../bin/polygon/stakingDualRewards.json';
import axios from "axios";
import * as Helpers from '../../helpers';
import { network } from '../../bin/token';
import TokenABI from '../../bin/tokenContract.ABI.json';
import pairContractABI from '../../bin/pairContractABI.json';


class Quickswap {
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
  public async calculateAPRValue(farm: any): Promise<string> {
    try {
      const quickNew = '0xb5c064f955d8e7f38fe0460c556a72987494ee17';
      const quickOld = '0x831753dd7087cac61ab5644b308642cc1c33dc13';
      const dQuick = '0xf28164a485b0b2c90639e47b0f377b4a438a16b1';
      const { masterchef, deposit_token, reward, chain_id, farmType } = farm;


      //reward token price (cake for pancake)
      let rewardTokenPrice = 0, depositTokenPrice = 0, tokenDecimals = 18;

      const tokenSymbol = await this.getSymbol(reward, chain_id);
      if (tokenSymbol === "QUICK") {
        // OLDQUICK
        rewardTokenPrice = await this.getTokenPriceUSD("DQUICK");
      } else {
        rewardTokenPrice = await this.getTokenPriceUSD(tokenSymbol);
        tokenDecimals = Number(await this.getDecimals(reward, chain_id));
      }
      let stakingRewardsABI = StakingRewardsABI;
      if (farmType === "quickswapdual") {
        stakingRewardsABI = StakingDualRewardsABI;
      }
      // total supply value
      if (deposit_token === quickNew) {
        depositTokenPrice = await this.getTokenPriceUSD('QUICK');
      } else {
        depositTokenPrice = await this.calPrice(deposit_token, chain_id);
      }
      // depositTokenPrice = 2;

      const stakingRewardContract: any = await Helpers.Web3Helper.callContract(chain_id, stakingRewardsABI, masterchef);

      const totalSupply: any = await stakingRewardContract.methods.totalSupply().call();

      let rewardRate: any = 0;
      if (farmType === "quickswapdual") {
        const rewardTokenA = await stakingRewardContract.methods.rewardsTokenA().call();
        const rewardTokenB = await stakingRewardContract.methods.rewardsTokenB().call();
        let rewardTokenASymbol = 'DQUICK';
        let rewardTokenADecimals = 18;
        let rewardTokenAPrice = rewardTokenPrice;
        // token is not dQuick
        if (rewardTokenA.toLowerCase() !== dQuick.toLowerCase()) {
          rewardTokenASymbol = await this.getSymbol(rewardTokenA, chain_id);
          rewardTokenADecimals = Number(await this.getDecimals(rewardTokenA, chain_id));
          rewardTokenAPrice = await this.getTokenPriceUSD(rewardTokenASymbol.toUpperCase());
        }

        const rewardTokenBSymbol = await this.getSymbol(rewardTokenB, chain_id);
        const rewardTokenBDecimals = Number(await this.getDecimals(rewardTokenB, chain_id));
        const rewardTokenBPrice = await this.getTokenPriceUSD(rewardTokenBSymbol.toUpperCase());

        const rewardPerTokenA = await stakingRewardContract.methods.rewardRateA().call();
        const rewardPerTokenB = await stakingRewardContract.methods.rewardRateB().call();
        
        // console.table({rewardTokenA, rewardPerTokenA, rewardPerTokenB, rewardTokenAPrice, rewardTokenBPrice,rewardTokenASymbol ,rewardTokenBSymbol});

        rewardRate = (Number(rewardPerTokenA) * rewardTokenAPrice) / (10 ** rewardTokenADecimals);
        rewardRate += (Number(rewardPerTokenB) * rewardTokenBPrice) / 10 ** rewardTokenBDecimals;
      } else {
        rewardRate = Number(await stakingRewardContract.methods.rewardRate().call());
        rewardRate = (rewardRate * rewardTokenPrice) / 10 ** Number(tokenDecimals);
      }

      //   '-------------------------')
      const rewardRateYearly = rewardRate * (24 * 3600 * 365);

      const liquidity = ((Number(totalSupply) * depositTokenPrice) / 10 ** 18);
      const apr: any = rewardRateYearly / liquidity * 100;

      console.table({ tokenSymbol, rewardTokenPrice, tokenDecimals, rewardRate, totalSupply, rewardRateYearly, depositTokenPrice, liquidity, apr });
      return apr.toFixed(4);
    } catch (err) {
      throw err;
    }
  }

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
  public async getDecimals(tokenAddress: any, chainId: number): Promise<string> {
    try {
      const contract: any = await Helpers.Web3Helper.callContract(chainId, TokenABI, tokenAddress);
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
  * call price decenteralize way
  * @param  {any} pairAddress
  * @param  {number} chainId
  * @returns Promise
  */
  public async calPrice(pairAddress: any, chainId: number): Promise<any> {
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
        const r = await this.getTokenPriceUSD(symbolSingle);
        return r;
      }

      if (tokenZero === '0x0000000000000000000000000000000000000000') {
        const symbolSingle = await this.getSymbol(pairAddress, chainId);
        // return await this.getTokenPriceUSD(symbolSingle);
        const r = await this.getTokenPriceUSD(symbolSingle);
        return r;
      } else {
        const tokenOne: any = await this.getTokenOne(pairAddress, chainId);
        const reserve: any = await this.getReserves(pairAddress, chainId);
        const symbolZero: any = await this.getSymbol(tokenZero, chainId);
        const symbolOne: any = await this.getSymbol(tokenOne, chainId);
        const decimalZero: any = await this.getDecimals(tokenZero, chainId);
        const decimalOne: any = await this.getDecimals(tokenOne, chainId);
        // fetching data from Api for token zero...
        const respTokenZero = await this.getTokenPriceUSD(symbolZero);
        if (respTokenZero) {
          priceTokenZero = respTokenZero * (reserve[0] / 10 ** decimalZero);
        }
        // fetching data from Api for token one...
        const respTokenOne = await this.getTokenPriceUSD(symbolOne);
        if (respTokenOne) {
          priceTokenOne = respTokenOne * (reserve[1] / 10 ** decimalOne);
        }
        // p0 = (reserve1/10**decimals1) / (reserve0/10**decimals0)
        price = (priceTokenOne + priceTokenZero);
        // price = priceTokenZero + priceTokenOne;
        // console.table({ price });

        return (price);
      }
    } catch (err) {
      throw (err);
    }
  }
}

export default new Quickswap();
