export interface Result {
    data: Data[];
}

export interface Data {
    year?: string;
    month?: string;
    count: number;
    cumulative?: number;
    genre?: string;
}

export interface AnalysisRequest {
    analysis: Analysis[];
    genre?: string;
    isCumulative?: boolean
}

export interface Analysis {
    year?: string;
    months?: string[];
}