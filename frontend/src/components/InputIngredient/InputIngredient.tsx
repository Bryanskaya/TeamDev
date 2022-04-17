import { Box, Button, useDisclosure } from '@chakra-ui/react'
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
} from '@chakra-ui/react'
import React from 'react'

import AddIcon from 'components/Icons/Add'
import Input from 'components/Input'
import { Ingredient as IngredientT } from 'types/Ingredient'

import styles from "./InputIngredient.module.scss"

export default function IngredientInput(propos) {
  const { isOpen, onOpen, onClose } = useDisclosure()
  var data: IngredientT = { title: '', amount: '' }

  async function put() { 
    await propos.putCallback(data)
    onClose()
  }

  return (
    <>
      <Button className={styles.add_btn} onClick={onOpen}>
        <Box> <AddIcon /> </Box>
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent className={styles.dark_bg}>
          <ModalHeader>Новый ингредиент</ModalHeader>
          <ModalCloseButton />
          <ModalBody className={styles.model_body}>
            <Box>
              <Input placeholder='Введите ингредиент'
                onInput={(e) => {data.title = e.currentTarget.value}}
              />
              <Input placeholder='Введите меру'
                onInput={(e) => {data.amount = e.currentTarget.value}}
              />
            </Box>
            <Button className={styles.ready_btn} onClick={put}>
              <AddIcon className={styles.img_btn} />
            </Button>
          </ModalBody>
          <ModalFooter/>
        </ModalContent>
      </Modal>
    </>
  )
}
