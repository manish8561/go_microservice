import BaseModel from "../../model/base.model";
// import request from "request";

import masterChefAbi from '../../bin/masterChefContractABI.json'
import masterChefHelper from "./masterChef.helper";
import farms from "../../model/schema/farms";
import { Responses } from "../../helpers";
import axios from "axios";


class farmModel extends BaseModel {
  constructor() {
    super();
  }


  public async getAprValue(response: any) {

    const ContarctAddess = '0x8a69E9780700c0B42825ED5F5dDf8ca0B6A3B6e0'
    const abi = masterChefAbi

    const farmData: any = await axios.get(`${process.env.FARM_API_URL}farm?page=1&limit=10&chain_id=4`);


    if (farmData.status === 200 && farmData.data.data.length > 0) {
      try {


        // if (farmData && farmData.length > 0) {
        // farmData.data.data.map(async (it: any) => {
        //   const { masterchef, deposit_token }: any = it
        //   const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token);
        //   console.log("call Apr", calApr)
        // })
        // }

        // await farms.insertMany(farmData.data.data, { ordered: false }).catch(err => {
        //   console.error(err);
        // })



        const farmList = await farms.find({});

        // if (farmList && farmList.length > 0) {
        //   farmList.map((it: any) => {
        //     const mC : any = it.masterchef
        //     console.log('HERE', it, i)
        //   })
        // }

        // console.log("FARMMMMMMMMMMMMMMM", farmList);
        // const calApr: any = await masterChefHelper.calculateAPRValue(abi, ContarctAddess);
        // const data: any = await new farms({ apr: calApr }).save();
        return farmList
      } catch (error) {
        throw Responses.error(response, { message: error });
      }

    }


    // }
  }


}

export default new farmModel();