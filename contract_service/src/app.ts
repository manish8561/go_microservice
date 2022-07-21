import cors from "cors";
import express from "express";
import helmet from "helmet";
import morgan from "morgan";
import * as rfs from "rotating-file-stream";
import * as fs from "fs";
import * as path from "path";
import { CronJob } from 'cron';
import swaggerUi from 'swagger-ui-express';
import * as http from "http";
import * as swaggerDocument from './swagger.json';
// import * as config from "../src/config/";
import { Controller } from "./interfaces";
import { errorMiddleware } from "./middlewares";
import * as Helpers from './helpers';

import farmsContact from './modules/contract/contract.model';

class App {
    public app: express.Application;
    public port: any;
    private server: http.Server;
    /**
     * @param  {Controller[]} controllers
     */
    constructor(controllers: Controller[]) {
        this.app = express();
        this.port = process.env.PORT ? process.env.PORT : 8082;
        //socket io code
        this.server = http.createServer(this.app);
        Helpers.MongodbHelper.connectMongoDB();
        this.initializeMiddlewares();
        this.initializeControllers(controllers);
        this.initializeErrorHandling();

    }
    /**
     * listen
     */
    public listen() {
        this.server.listen(this.port, () => {
            console.log(
                `-- App listening on the port ${this.port
                }`
            );
        });
        return this.server;
    }
    /**
     * initializeMiddlewares
     */
    private initializeMiddlewares() {
        this.app.use(express.json());
        this.app.use(express.urlencoded());
        this.app.use(cors());
        this.app.use(helmet());
        // setup the logger
        this.saveLogs();
        //setting up swagger
        // this.useSwagger();
        this.startCron();
    }
    /**
     * swagger
     */
    private useSwagger() {
        const enable_swagger = process.env.ENABLE_SWAGGER == 'true' ? true : false;
        if (enable_swagger) {
            this.app.use('/explorer', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
            // this.app.use('/explorer', swaggerUi.serve);
            // this.app.get('/explorer', swaggerUi.setup(swaggerDocument));
        }
    }
    /**
     * initialize controller
     * @param  {Controller[]} controllers
     */
    private initializeControllers(controllers: Controller[]) {
        controllers.forEach(controller => {
            this.app.use("/api/contract_service", controller.router);
        });
        this.app.get("/api/contract_service/status", (req, res) => {
            return res.status(200).send({ status: "success" });
        });
    }
    /**
     * initialize error handling
     */
    private initializeErrorHandling() {
        this.app.use(errorMiddleware);
    }
    /**
     * save logs
     */
    private saveLogs() {
        console.log('\n inside savelogs ------ ');
        const logDirectory = path.join(__dirname, "log");
        // ensure log directory exists
        const exists = fs.existsSync(logDirectory) || fs.mkdirSync(logDirectory);
        if (exists) {
            const accessLogStream = rfs.createStream("access.log", {
                // size: "10M", // rotate every 10 MegaBytes written
                interval: "1d", // rotate daily
                path: logDirectory,
                // compress: "gzip" // compress rotated files
            });

            // setup the logger
            this.app.use(morgan("combined", { stream: accessLogStream }));
        }
    }
    /**
     * start cron job
     */
    private startCron() {
        const job = new CronJob('*/30 * * * * *', async () => {
            console.log('You will see this message every 1:00 am' + new Date());
            await farmsContact.getFarmsValue();
        }, null, true, 'Europe/London');
        job.start();
    }
}

export default App;
