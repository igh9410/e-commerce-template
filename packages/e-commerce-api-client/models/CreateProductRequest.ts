/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Request message for creating a product
 */
export type CreateProductRequest = {
    name: string;
    description?: string;
    category?: string;
    price: string;
    status?: CreateProductRequest.status;
    tags?: Array<string>;
};
export namespace CreateProductRequest {
    export enum status {
        ACTIVE = 'ACTIVE',
        INACTIVE = 'INACTIVE',
    }
}

