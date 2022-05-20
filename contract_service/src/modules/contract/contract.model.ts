import BaseModel from "../../model/base.model";
import masterChefHelper from "./masterChef.helper";
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
      const farmLength: any = await Farm.countDocuments({ status: 'active' });
      let arr: any = [];
      if (farmLength > 0) {
        for (let i = 0; i < farmLength; i += 20) {
          const farmData: any = await Farm.find({ status: 'active' }).skip(i).limit(20);

          for (let it of farmData) {
            arr = [...arr, this.getFarm(it)];
          }
        }
        await Promise.all(arr);
        return { success: 'updated Successfully' }
      }
      else {
        return { success: 'No data found' }
      }
    } catch (error) {
      console.log("err", error)
    }
  }
  /**
   * get farm details
   * @param  {any} it
   * @returns Promise
   */
  private async getFarm(it: any): Promise<void> {
    try {
      const { chain_id, pid, masterchef, deposit_token, token_type, address }: any = it;
      const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token, chain_id, pid);
      const calApy: any = await masterChefHelper.calculateAPY(calApr);
      const calTvl: any = await masterChefHelper.calculateTVLValue(deposit_token, address, token_type, chain_id);
      it.daily_apr = calApr
      it.tvl_staked = calTvl
      it.daily_apy = calApy
      it.save()
    } catch (error) {
      console.log(error, '-----------------')
    }
  }
}

export default new farmModel();