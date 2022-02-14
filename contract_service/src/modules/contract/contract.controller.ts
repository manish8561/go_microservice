import { Response, Router } from "express";
import * as Interfaces from "interfaces";
import { Responses } from "../../helpers";
// import { requestDecrypt } from "../../middlewares";
import Contract from "./contract.model";
import masterChefAbi from '../../bin/masterChefContractABI.json'
import masterChefHelper from "./masterChef.helper";

class ContractController implements Interfaces.Controller {
    public path = '/Contract';
    public router = Router();
    constructor() {
        this.initializeRoutes();
    }

    private async initializeRoutes() {
        this.router
            .all(`${this.path}/*`)

            // .get(
            //     this.path + "/getTotalPairs",
            //     this.getTotalPairs
            // )
            // .get(
            //     this.path + "/getTotalLiquidity",
            //     this.getTotalLiquidity
            // )
            // .get(
            //     this.path + "/getTotalVolume",
            //     this.getTotalVolume
            // )
            .get(
                this.path + "/getapr",
                this.getAprValue
            );
    }


    // Get AprValue of lp
    private async getAprValue(req: any, response: Response) {
        if (req.body) {
            // const chainId = 1280, factory = "0x2A6d5103E4312E487413E476B7aF91FD67be7627";
            const ContarctAddess = '0x8a69E9780700c0B42825ED5F5dDf8ca0B6A3B6e0'
            const abi = masterChefAbi
            try {
                const resp: string = await masterChefHelper.calculateAPRValue(abi, ContarctAddess,);
                return Responses.success(response, { resp });
            } catch (error) {
                return Responses.error(response, { message: error });
            }
        }
    }
}

export default ContractController;