import React from "react";
import theme from "styles/extendTheme";

import { Box, Link } from "@chakra-ui/react";
import { NavigateFunction } from "react-router-dom";
import { CookieSetOptions } from "universal-cookie";

import { Account } from "types/Account";
import { Create as CreateQuery } from "postAPI/accounts/Create";
import { Login as LoginQuery } from "postAPI/accounts/Login";

import Input from "components/Input";
import RoundButton from "components/RoundButton";

import styles from "./SignupPage.module.scss";

type SignUpProps = {
    navigate: NavigateFunction
    cookie: {
        token?: string;
        role?: string;
        login?: string;
    }
    setCookie: (name: "token" | "role" | "login", value: any, options?: CookieSetOptions | undefined) => void
}


class SignUpPage extends React.Component<SignUpProps> {
    acc: Account = {login: ""}
    repPassword: string = ""

    setLogin(val: string) {
        this.acc.login = val
    }
    setPassword(val: string) {
        this.acc.password = val
    }
    setRepPassword(val: string) {
        this.repPassword = val
    }

    highlightNotMatch() {
        let node1 = document.getElementsByName("password")[0]
        let node2 = document.getElementsByName("rep-password")[0]

        if (node1.parentElement && node2.parentElement) {
            node1.parentElement.style.borderColor = theme.colors["title"]
            node2.parentElement.style.borderColor = theme.colors["title"]
        }

        var title = document.getElementById("undertitle")
        if (title)
            title.innerText = "Пароли не совпадают!"
    }

    async submit(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        if (this.acc.password !== this.repPassword)
            return this.highlightNotMatch()

        e.currentTarget.disabled = true
        var data = await CreateQuery(this.acc)
        if (data.status === 200) {
            await LoginQuery(this.acc, this.props.setCookie)
            window.location.href = '/';
        } else {
            e.currentTarget.disabled = false
            var title = document.getElementById("undertitle")
            if (title)
                title.innerText = "Ошибка создания аккаунта!"
        };
    }

    render() {
        return <Box className={styles.login_page}>
            <Box className={styles.input_div}>
                <Input name="login" placeholder="Введите логин" 
                onInput={event => this.setLogin(event.currentTarget.value)}/>
                <Input name="password" type="password" placeholder="Введите пароль"
                onInput={event => this.setPassword(event.currentTarget.value)}/>
                <Input name="rep-password" type="password" placeholder="Повторите пароль"
                onInput={event => this.setRepPassword(event.currentTarget.value)}/>
            </Box>

            <Box className={styles.input_div}>
                <RoundButton type="submit" onClick={event => this.submit(event)}>
                    Создать аккаунт
                </RoundButton>
                <Link href="/auth/signin">Войти</Link>
            </Box>
        </Box>
    }
}

export default SignUpPage;