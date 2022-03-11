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
  public async getAprValue(response: any) {
    try {
      const farmLength: any = await farms.count({});
      if (farmLength > 0) {
        for (let i = 0; i <= farmLength; i += 20) {
          console.log(`************ ${i} ***************`)
          const farmData: any = await farms.find().skip(i).limit(20);
          for (let it of farmData) {
            console.log(`************ time ***************`)
            const { masterchef, deposit_token , token_type , address }: any = it;
            const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token);
            const calTvl: any = await masterChefHelper.calculateTVLValue(deposit_token , address);
            it.daily_apr = calApr
            it.save()
          }
        }
        return {}
      }
    }
    catch (error) {
      throw Responses.error(response, { message: error });
    }
  }
}

export default new farmModel();