import React from "react";

import { Box, Link } from "@chakra-ui/react";
import { NavigateFunction } from "react-router-dom";
import { CookieSetOptions } from "universal-cookie";

import Input from "components/Input";
import RoundButton from "components/RoundButton";

import { Account } from "types/Account"
import { Login as LoginQuery } from "postAPI/accounts/Login";

import styles from "./LoginPage.module.scss";

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
        return <Box className={styles.login_page}>
            <Box className={styles.input_div}>
                <Input name="login" placeholder="Введите логин"
                onInput={event => this.setLogin(event.currentTarget.value)}/>
                <Input name="password" type="password" placeholder="Введите пароль"
                onInput={event => this.setPassword(event.currentTarget.value)}/>
            </Box>

            <Box className={styles.button_div}>
                <RoundButton onClick={ (event) => this.submit(event) }> Войти </RoundButton>
                <Link href="/auth/signup">Зарегистрироваться</Link>
                <Link href="/">Назад</Link>
            </Box>
        </Box>
    }
}

export default LoginPage;
