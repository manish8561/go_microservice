import BaseModel from "../../model/base.model";
import masterChefHelper from "./masterChef.helper";
import farms from "../../model/schema/farms";

class farmModel extends BaseModel {
  constructor() {
    super();
  }
  public async getFarmsValue() {
    try {
      // const addData = await axios.get('https://ws-stage.autocompound.com/api/api/farm_service/farm?page=1&limit=10&chain_id=4');
      // farms.insertMany(addData.data.data)
      const farmLength: any = await farms.count({});
      if (farmLength > 0) {
        for (let i = 0; i < farmLength; i += 20) {
          const farmData: any = await farms.find({ status: 'active' }).skip(i).limit(20);
          for (let it of farmData) {
            const { masterchef, deposit_token, token_type, address }: any = it;
            const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token);
            console.log('hi', calApr)
            const calApy: any = await masterChefHelper.calculateAPY(calApr);
            const calTvl: any = await masterChefHelper.calculateTVLValue(deposit_token, address, token_type);
            it.daily_apr = calApr
            it.tvl_staked = calTvl
            it.daily_apy = calApy
            it.save()
          }
        }
        return { success: 'updated Successfully' }
      }
      else {
        return { success: 'No data found' }
      }
    }
    catch (error) {
      console.log("err", error)
    }
  }
}

export default new farmModel();