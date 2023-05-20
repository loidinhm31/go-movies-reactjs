export interface Result {
    data: Data[];
}

export interface Data {
    year?: string;
    month?: string;
    count: number;
    cumulative?: number;
    name?: string;
    type_code?: string;
}

export interface AnalysisRequest {
    analysis: Analysis[];
    name?: string;
    type_code?: string;
    isCumulative?: boolean
}

export interface Analysis {
    year?: string;
    months?: string[];
}