export class ClientError {
    message: any;
    errorCode: number;
    httpStatusCode: number;

    constructor({
        errorCode,
        httpStatusCode,
        message,
    }: {
        message: string;
        errorCode: number;
        httpStatusCode: number;
    }) {
        this.message = message;
        this.errorCode = errorCode;
        this.httpStatusCode = httpStatusCode;
    }

    toString() {
        return JSON.stringify(this);
    }
}
