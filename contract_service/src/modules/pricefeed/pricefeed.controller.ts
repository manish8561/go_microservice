import { Request, Response, Router } from "express";

import * as Interfaces from "interfaces";
import { Responses } from "../../helpers";

import PriceFeed from "./pricefeed.model";


class PriceFeedController implements Interfaces.Controller {
  public path = '/PriceFeed';
  public router = Router();
  constructor() {
    this.initializeRoutes();
  }

  private async initializeRoutes() {
    this.router
      .all(`${this.path}/*`)
      
      .get(
        this.path + '/getPrice',
        this.getPrice
      );
  }
   
  /**
   * get price
   * @param  {any} req
   * @param  {Response} response
   */
  private async getPrice(req: any, response: Response) {
    try {
        const {symbol} = req.query;
        console.log(symbol, 'before')
      const res: any = await PriceFeed.getPrice('moon-rabbit');
      return Responses.success(response, {symbol:"RS", price: (res['moon-rabbit'].usd*600)});
    } catch (error) {
      return Responses.error(response, { message: error });
    }
  }
}

export default PriceFeedController;