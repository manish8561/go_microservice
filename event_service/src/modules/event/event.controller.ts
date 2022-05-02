import { Response, Router } from "express";
import * as Interfaces from "interfaces";
import { Responses } from "../../helpers";
// import { requestDecrypt } from "../../middlewares";

import EventModel from './event.model'

class EventController implements Interfaces.Controller {
    public path = '/event';
    public router = Router();
    constructor() {
        this.initializeRoutes();
    }

    private async initializeRoutes() {
        this.router
            .all(`${this.path}/*`)
            .get(
                this.path + "/getapr",
                this.calApy
            );
    }

    private async calApy(req: any, response: Response) {
        try {
            const res: any = await EventModel.getFarmsValue();
            return Responses.success(response, { message: res });;
        } catch (error) {
            return Responses.error(response, { message: error });
        }
    }
}

export default EventController;