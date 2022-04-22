import { Recipe } from "types/Recipe";
import { Account } from "types/Account";

export const backUrl = "http://localhost:8000";

export type AllRecipeResp = {
    status: number,
    content: Recipe[] | string
}

export type AllCategoriesResp = {
    status: number,
    content: string[]
}

export type AllUsersResp = {
    status: number,
    content: Account[]
}
