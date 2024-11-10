/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
/**
 * Product message with tags as repeated field
 */
export type Product = {
    id?: string;
    name?: string;
    description?: string;
    category?: string;
    price?: string;
    status?: Product.status;
    tags?: Array<string>;
};
export namespace Product {
    export enum status {
        ACTIVE = 'ACTIVE',
        INACTIVE = 'INACTIVE',
    }
}

