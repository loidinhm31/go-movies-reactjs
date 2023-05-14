export type PageType<T> = {
    size?: number,
    page?: number,
    total_elements?: number,
    total_pages?: number,
    content?: T[],
    sort?: SortType
}

export type SortType = {
    orders: OrderType[]
}

export type OrderType = {
    property: string,
    direction: Direction,
}

export enum Direction {
    ASC = "asc", DESC = "desc"
}