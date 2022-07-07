import StakingRewardsABI from '../../bin/polygon/stakingRewards.json';
import StakingDualRewardsABI from '../../bin/polygon/stakingDualRewards.json';
import axios from "axios";
import * as Helpers from '../../helpers';
import { network } from '../../bin/token';

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
  public async calculateAPRValue(stakingRewards: string, lp: string, chainId: number, pid: number, farmType: string): Promise<string> {
    try {
      //reward token price (cake for pancake)
      let acPrice = 0;
      acPrice = await this.getTokenPriceUSD('AC');

      let stakingRewardsABI = StakingRewardsABI;
      if (farmType === "quickswapdual") {
        stakingRewardsABI = StakingDualRewardsABI;
      }

      const stakingRewardContract: any = await Helpers.Web3Helper.callContract(chainId, stakingRewardsABI, stakingRewards);

      const totalSupply: any = await stakingRewardContract.methods.totalSupply().call();
      let cakePerBlock: any = 0;
      if (farmType === "quickswapdual") {
        const rewardPerTokenA = await stakingRewardContract.methods.rewardRateA().call();
        const rewardPerTokenB = await stakingRewardContract.methods.rewardRateB().call();
        cakePerBlock = Number(rewardPerTokenA) + Number(rewardPerTokenB);
      } else {
        cakePerBlock = Number(await stakingRewardContract.methods.rewardRate().call());
      }
      const blockMined = network[chainId].blockMined;
      // console.log(
      //   'cake per block', cakePerBlock,
      //   'total supply', totalSupply,
      //   '-------------------------');
      //since it is in cake value
      // const accCakePerShare = poolInfo.accCakePerShare / (10 ** 18);
      // const apr: any = ((accCakePerShare / totalAllcationPoint) * ((cakePerBlock / 10 ** 18) * blockMined * 100 * acPrice)) / liquidity;
      const apr: any = ((cakePerBlock / totalSupply) * (blockMined * 100 * acPrice));
      return apr.toFixed(4);
    } catch (err) {
      throw err;
    }
  }
}

export default new Quickswap();
