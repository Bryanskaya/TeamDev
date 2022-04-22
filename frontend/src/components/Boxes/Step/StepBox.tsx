import { Box, HStack, Text, VStack, Button } from "@chakra-ui/react";
import React from "react";

import Input from "components/Input";
import {Step as StepT} from "types/Step";

import PenIcon from "components/Icons/Pen";
import DeleteIcon from "components/Icons/DeleteSmall";
import OkIcon from "components/Icons/Ok";

import PatchStep from "postAPI/steps/Patch";

import styles from "./StepBox.module.scss";


async function updateChanges(step: StepT) {
    await PatchStep(step);
}


interface StepProps extends StepT{
    delStepCallback?: (num: number) => Promise<void>
}


const StepBox: React.FC<StepProps> = (props) => {
    const [active, change] = React.useState(false);
    const [title, changeTitle] = React.useState(props.title);
    const [description, changeDescr] = React.useState(props.description);
    const textareaRef = React.useRef<HTMLTextAreaElement>(null);

    React.useLayoutEffect(() => {
        if (textareaRef.current) {
            textareaRef.current.style.height = "inherit";
            textareaRef.current.style.height = `${Math.max(
            textareaRef.current.scrollHeight,
            20
            )}px`;
        }
    }, [description, active]);

    return (
        <Box className={styles.main_box}>
            <VStack>
                <HStack className={styles.top_box}>
                    {active && <Input className={styles.title}
                        value={title} placeholder='Введите название' 
                        onInput={(e) => {changeTitle(e.currentTarget.value)}}>
                    </Input>}

                    {!active && <Text>
                        {title}
                    </Text>}

                    {props.delStepCallback && <HStack> 
                        <Button className={styles.btn} onClick={() => 
                            {
                                active && updateChanges({recipe: props.recipe, num: props.num, title: title, description: description}); 
                                change(!active)}
                            }>

                            {!active && <PenIcon />}
                            {active && <OkIcon/>}
                        </Button> 

                        <Button className={styles.btn}
                            onClick={() => props.delStepCallback && props.delStepCallback(props.num)}>
                            <DeleteIcon /> 
                        </Button>
                    </HStack>}
                </HStack>

                {active && <textarea className={styles.description_input} ref={textareaRef}
                    value={description} placeholder="Введите описание"
                    onChange={(e) => {changeDescr(e.currentTarget.value)}}></textarea>}

                {!active && <Box className={styles.description_box}> 
                    <Text>
                        {description}
                    </Text>
                </Box>}
            </VStack>
        </Box>
    )
}

export default StepBox;