import {PageType} from "./page";

export type SearchRequest = {
    filters?: FieldData[],
    page_request?: PageType<unknown>
}

export type FieldData = {
    operator?: string,
    field: string,
    def?: TypeValue,
}

export type TypeValue = {
    type: string,
    values: string[],
}

