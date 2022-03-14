import BaseModel from "../../model/base.model";
import masterChefHelper from "./masterChef.helper";
import farms from "../../model/schema/farms";
import { Responses } from "../../helpers";
// import axios from "axios";
// import { initParams } from "request-promise";


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
        for (let i = 0; i <= farmLength; i += 20) {
          const farmData: any = await farms.find().skip(i).limit(20);
          for (let it of farmData) {
            const { masterchef, deposit_token , token_type , address }: any = it;
            const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token);
            const calTvl: any = await masterChefHelper.calculateTVLValue(deposit_token , address);
            const calApy: any = await masterChefHelper.calculateAPY(masterchef, deposit_token);
            it.daily_apr = calApr
            it.tvl_staked = calTvl
            it.daily_apy = calApy
            it.save()
          }
        }
        return {}
      }
    }
    catch (error) {
      console.log("err", error)
      // throw Responses.error(response, { message: error });
    }
  }
}

export default new farmModel();