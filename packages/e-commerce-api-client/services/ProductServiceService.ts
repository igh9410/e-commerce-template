/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CreateProductRequest } from '../models/CreateProductRequest';
import type { CreateProductResponse } from '../models/CreateProductResponse';
import type { Status } from '../models/Status';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class ProductServiceService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * POST /api/v1/products
     * @param requestBody
     * @returns CreateProductResponse OK
     * @returns Status Default error response
     * @throws ApiError
     */
    public productServiceCreateProduct(
        requestBody: CreateProductRequest,
    ): CancelablePromise<CreateProductResponse | Status> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/api/v1/products',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
}
