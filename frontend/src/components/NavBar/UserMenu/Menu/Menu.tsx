import { Heading, majorScale, Menu as EvergreenMenu, Pane, Text } from 'evergreen-ui'
import * as React from 'react'
import { Link } from 'react-router-dom'
import { User } from '../../../../lib/auth'

declare interface MenuProps {
  user: User
}

export default class Menu extends React.Component<MenuProps> {
  public render() {
    const props = this.props

    return (
      <EvergreenMenu>
        <Pane
          padding={majorScale(2)}
          maxWidth={majorScale(30)}
          // backgroundColor={defaultTheme.scales.neutral.N1A}
        >
          <Heading is="h3">{props.user.firstName}</Heading>
          <Pane>
            <Text>{props.user.email}</Text>
          </Pane>
        </Pane>

        <EvergreenMenu.Divider />

        <EvergreenMenu.Group>
          <EvergreenMenu.Item is={Link} href="/logout" rel="external">
            Logout
          </EvergreenMenu.Item>
        </EvergreenMenu.Group>
      </EvergreenMenu>
    )
  }
}
