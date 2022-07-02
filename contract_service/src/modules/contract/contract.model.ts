import BaseModel from "../../model/base.model";
import masterChefHelper from "./masterChef.helper";
import quickswapHelper from "./quickswap.helper";
import Farm from "../../model/schema/farms";

class farmModel extends BaseModel {
  constructor() {
    super();
  }
  /**
   * get farm value
   */
  public async getFarmsValue() {
    try {
      const now = Date.now();
      const farmLength: any = await Farm.countDocuments({ status: 'active' });
      let arr: any = [];
      if (farmLength > 0) {
        console.log((Date.now() - now) / 1000, 'after db fetch---------------------');
        for (let i = 0; i < farmLength; i += 20) {
          const farmData: any = await Farm.find({ status: 'active' }).skip(i).limit(20);

          for (let it of farmData) {
            arr = [...arr, this.getFarm(it)];
          }
        }

        await Promise.all(arr);
        console.log((Date.now() - now) / 1000, 'after  loop---------------------');

        return { success: 'updated Successfully' };
      }
      else {
        return { success: 'No data found' };
      }
    } catch (error) {
      console.log("err", error);
    }
  }
  /**
   * get farm details
   * @param  {any} it
   * @returns Promise
   */
  private async getFarm(it: any): Promise<void> {
    try {
      const now = Date.now();

      const { chain_id, pid, masterchef, deposit_token, token_type, address, ac_token, farmType }: any = it;

      let calApr: any = 0, calApy: any = 0;

      let calTvl: any = await masterChefHelper.calculateTVLValue(deposit_token, address, token_type, chain_id);

      let acPerBlock: Number = await masterChefHelper.getTokenPerBlock(ac_token, address, chain_id);
      if (farmType === "quickswap" || farmType === "quickswapdual") {
        calApr = await quickswapHelper.calculateAPRValue(masterchef, deposit_token, chain_id, pid, farmType);
      } else {
        calApr = await masterChefHelper.calculateAPRValue(masterchef, deposit_token, chain_id, pid);

      }

      calApy = await masterChefHelper.calculateAPY(calApr);
      it.daily_apr = Number(calApr);
      it.daily_apy = Number(calApy);
      it.tvl_staked = calTvl;
      it.token_per_block = acPerBlock;
      await it.save();
      console.log((Date.now() - now) / 1000, 'inside loop---------------------');
    } catch (error) {
      console.log(it.address, ':error strategy');
      console.log(error, '-----------------get farm function--------');
    }
  }
}

export default new farmModel();