import React from "react";
import { NavigateFunction, Params } from "react-router-dom";
import { Box, Button, HStack, Image, Text, VStack } from "@chakra-ui/react";

import GetRecipe from "postAPI/recipes/Get"
import GetIngredient from "postAPI/ingredients/Get";
import GetSteps from "postAPI/steps/Get";
import PushStep from "postAPI/steps/Push";
import GetLikes from "postAPI/likes/Get";
import DeleteLike from "postAPI/likes/Delete";
import AddLike from "postAPI/likes/Add";
import GetIsLiked from "postAPI/likes/GetLiked";
import DeleteIngredient from "postAPI/ingredients/Delete";
import DeleteStep from "postAPI/steps/Delete";
import PutIngredient from "postAPI/ingredients/Post";
import DeleteRecipe from "postAPI/recipes/Delete";

import {Recipe as RecipeT} from "types/Recipe"
import {Ingredient as IngredientT} from "types/Ingredient";
import {Step as StepT} from "types/Step";

import DeleteIcon from "components/Icons/Delete";
import EmptyLike from "components/Icons/EmptyLike";
import FullLike from "components/Icons/FullLike";
import AddIcon from "components/Icons/Add";

import Ingredient from "components/IngredientBox";
import IngredientModel from "components/InputIngredient";
import ClockBox from "components/Boxes/Clock";
import FullLikeBox from "components/Boxes/FullLike";
import PortionBox from "components/Boxes/Portion";

import styles from "./RecipeInfoPage.module.scss";
import StepBox from "components/Boxes/Step";


type State = {
    recipe?: RecipeT,
    liked: boolean,
    ingredients: Array<IngredientT>,
    steps: Array<StepT>,
    likes: number,

    isAuthor: boolean,
}

type RecipeInfoParams = {
    match: Readonly<Params<string>>
    navigate: NavigateFunction
    cookie: {
        role?: string;
        login?: string;
    }
}

class RecipeInfoPage extends React.Component<RecipeInfoParams, State> {
    id: number;
    
    constructor(props) {
      super(props);
      this.id = Number(this.props.match.id)
      this.state = {
          recipe:undefined, 
          liked:false, 
          ingredients:[], 
          steps:[],
          likes:0,
          isAuthor: false
        }
    }

    async deleteRecipe() {
        var data = await DeleteRecipe(this.id);
        if (data.status === 200) {
            window.location.href = '/';
        }
    }

    async deleteLike() {
        var data = await DeleteLike(this.id);
        if (data.status === 200) {
            this.setState({liked:false});
            this.getLikes();
        }
    }

    async addLike() {
        var data = await AddLike(this.id);
        if (data.status === 200) {
            this.setState({liked:true});
            this.getLikes();
        }
    }

    async getLikes() {
        var data = await GetLikes(this.id)
        if (data.status === 200) {
            this.setState({likes: data.content})
        }
    }

    async getIsLiked() {
        var data = await GetIsLiked(this.id)
        if (data.status === 200) {
            this.setState({liked: data.content})
        }
    }

    async createStep() {
        let step: StepT = {
            description:    '',
            title: '',
            num: 0,
            recipe: this.id
        }

        var data = await PushStep(this.id, step)
        if (data.status === 200) {
            var temp = this.state.steps
            temp.push(data.content)
            this.setState({steps: temp})
        }
    }

    async putIngredient(data: IngredientT) {
        var res = await PutIngredient(this.id, data)
        if (res.status === 201) {
            var ingArr = this.state.ingredients
            ingArr.push(data)
            this.setState({ingredients: ingArr})
        }
    }

    async deleteIngredient(title: string) {
        var data = await DeleteIngredient(this.id, title)
        if (data.status === 200) {
            var ingArr = this.state.ingredients
            ingArr = ingArr.filter(item => item.title !== title)
            this.setState({ingredients: ingArr})
        }
    }

    async deleteStep(num: number) {
        var data = await DeleteStep(this.id, num)
        if (data.status === 200) {
            var temp = this.state.steps
            temp = temp.filter(item => item.num !== num)
            this.setState({steps: temp})
        }
    }

    componentDidMount() {
        GetRecipe(this.id).then(data => {
            if (data.status === 200) {
                this.setState({recipe: data.content})
                this.setState({isAuthor: this.state.recipe?.author === this.props.cookie.login})

                var elem = document.getElementById("title")
                if (elem && this.state.recipe)
                    elem.innerText = this.state.recipe.title

                elem = document.getElementById("author")
                if (elem && this.state.recipe)
                    elem.innerText = this.state.recipe.author
            }
        });

        GetIngredient(this.id).then(data => {
            if (data.status === 200) {
                this.setState({ingredients: data.content})
            }
        });

        GetSteps(this.id).then(data => {
            if (data.status === 200) {
                this.setState({steps: data.content})
            }
        });

        this.getLikes();
        if (this.props.cookie.login) this.getIsLiked();
    }
  
    render() {
        return(
            <VStack className={styles.flex_style}>
                <Box className={styles.flex_style}>
                    <HStack className={styles.img_block}>
                        <Image src={this.state.recipe?.pic_url} />

                        <VStack>
                            <Button onClick={() => this.deleteRecipe()}>
                                <Box className={styles.delete_box}> 
                                    {this.state.isAuthor && <DeleteIcon className={styles.delete_icon} /> }
                                </Box>
                            </Button>
                            
                            <Button onClick={() => {                                
                                (this.state.liked && this.deleteLike()) 
                                    || this.addLike()
                            }
                            }>

                                <Box className={styles.like_box}>
                                    {this.props.cookie.login 
                                    &&  ((this.state.liked && <FullLike className={styles.like_icon}/> )
                                        || <EmptyLike className={styles.like_icon} />)
                                    }
                                </Box>
                            </Button>
                        </VStack>
                    </HStack>

                    <Box className={styles.info_block}>
                        <HStack className={styles.items_block}>
                            <ClockBox duration={this.state.recipe?.duration}/>
                            <FullLikeBox likesNum={this.state.likes}/>
                            <PortionBox portionNum={this.state.recipe?.portion_num}/>
                        </HStack>

                        <Text>
                            {this.state.recipe?.description}
                        </Text>

                        <Box>
                            <HStack className={styles.ingredient_box}>
                                <Text>
                                    Ингредиенты
                                </Text>

                                {this.state.isAuthor &&
                                 <IngredientModel putCallback={(data: IngredientT) => this.putIngredient(data)}/>
                                }
                            </HStack>

                            <Box className={styles.ings_box}> 
                                <Box>
                                    {this.state.ingredients.map(item =>
                                        <Ingredient {...item} key={item.title} 
                                        delCallback={
                                            this.state.isAuthor 
                                            ? (title) => this.deleteIngredient(title) 
                                            : undefined
                                        } 
                                        />
                                    )}
                                </Box>
                            </Box>
                        </Box>
                    </Box>
                </Box>

                <Box className={styles.step_box}>
                    <HStack>
                        <Text>
                            Шаги приготовления
                        </Text>

                        <Button onClick={() => this.createStep()}>
                            {this.state.isAuthor &&
                                <Box> <AddIcon /> </Box>}
                        </Button>
                    </HStack>

                    <Box>
                        {this.state.steps.map(item =>
                            <StepBox {...item} key={item.num}
                            delStepCallback={
                                this.state.isAuthor 
                                ? (num) => this.deleteStep(num)
                                : undefined
                            }
                            />
                        )}
                    </Box>
                </Box>
            </VStack> 
      )
    }  
}

export default RecipeInfoPage;

