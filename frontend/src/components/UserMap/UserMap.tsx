import { Box } from "@chakra-ui/react";
import React from "react";
import UserCard from "../UserCard";
import { AllUsersResp } from "postAPI";

import styles from "./UserMap.module.scss";

interface UserBoxProps {
  searchQuery?: string;
  getCall: (title?: string) => Promise<AllUsersResp>;
}

type State = {
  postContent?: any;
};

class UserMap extends React.Component<UserBoxProps, State> {
  constructor(props) {
    super(props);
    this.state = {
      postContent: [],
    };
  }

  async getAll() {
    var data = await this.props.getCall(this.props.searchQuery);
    if (data.status === 200) this.setState({ postContent: data.content });
  }

  componentDidMount() {
    this.getAll();
  }

  componentDidUpdate(prevProps) {
    if (this.props.searchQuery !== prevProps.searchQuery) {
      this.getAll();
    }
  }

  render() {
    return (
      <Box className={styles.map_box}>
        {this.state.postContent.map((item) => (
          <UserCard {...item} key={item.login} role={item.role} />
        ))}
      </Box>
    );
  }
}

export default React.memo(UserMap);
