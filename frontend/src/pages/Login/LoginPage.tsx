import React from "react";

import { Box, Link, FormControl } from "@chakra-ui/react";
import { NavigateFunction } from "react-router-dom";
import { CookieSetOptions } from "universal-cookie";

import Input from "components/Input";
import RoundButton from "components/RoundButton";

import { Account } from "types/Account"
import { Login as LoginQuery } from "postAPI/accounts/Login";

type LoginProps = {
    navigate: NavigateFunction
    cookie: {
        token?: string;
        role?: string;
        login?: string;
    }
    setCookie: (name: "token" | "role" | "login", value: any, options?: CookieSetOptions | undefined) => void
}


class LoginPage extends React.Component<LoginProps> {
    acc: Account = {login: ""}

    setLogin(val: string) {
        this.acc.login = val
    }
    setPassword(val: string) {
        this.acc.password = val
    }

    submit(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        e.currentTarget.disabled = true
        LoginQuery(this.acc, this.props.setCookie).then(data => {
            if (data.status === 200) {
                window.location.href = '/';
            } else {
                e.currentTarget.disabled = false
                var title = document.getElementById("undertitle")
                if (title)
                    title.innerText = "Ошибка авторизации!"
            }
        });
    }

    render() {
        return <Box
            display="flex" width="70%"
            flexDir="column"
            alignItems="stretch"
            justifyContent="space-around"
            rowGap="70px"
        >
            <Box d="flex" flexDirection="column" rowGap="35px">
                <FormControl isRequired>
                    <Input name="login" width="100%" placeholder="Введите логин" height="40px"
                    onInput={event => this.setLogin(event.currentTarget.value)}/>
                </FormControl>
                <FormControl isRequired>
                    <Input name="password" type="password" width="100%" placeholder="Введите пароль" height="40px"
                    onInput={event => this.setPassword(event.currentTarget.value)}/>
                </FormControl>
            </Box>

            <Box d="flex" flexDirection="column" rowGap="15px" alignItems="center">
                <RoundButton onClick={ (event) => this.submit(event) }
                    width="100%" height="40px" bg="button" textColor="bg"
                >
                    Войти
                </RoundButton>
                <Link href="/auth/signup">Зарегистрироваться</Link>
                <Link href="/">Назад</Link>
            </Box>
        </Box>
    }
}

export default LoginPage;
