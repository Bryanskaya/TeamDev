import { Recipe } from "types/Recipe";

export const backUrl = "http://localhost:8000";

export type AllRecipeResp = {
    status: number,
    content: Recipe[] | string
}
