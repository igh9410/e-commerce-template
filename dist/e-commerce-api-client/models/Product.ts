/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export type Product = {
    id?: string;
    name?: string;
    description?: string;
    category?: string;
    price?: string;
    status?: Product.status;
};
export namespace Product {
    export enum status {
        ACTIVE = 'ACTIVE',
        INACTIVE = 'INACTIVE',
    }
}

