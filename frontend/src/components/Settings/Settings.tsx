import { Card, Heading, majorScale, Pane, Button } from 'evergreen-ui'
import * as React from 'react'
import PlaidLink from 'react-plaid-link'
import { User } from '../../clients'

export interface SettingsProps {
  user: User
  plaidEnv: string
  plaidPublicKey: string
  handleOnLinkExit: (err: Error | undefined) => void
  handleOnLinkSuccess: (token: string, metadata: any) => void
}

export default class Settings extends React.Component<SettingsProps> {
  public render() {
    const props = this.props

    return (
      <Pane display="flex" justifyContent="center">
        <Card
          width={majorScale(80)}
          height="auto"
          elevation={0}
          backgroundColor="white"
          marginY={majorScale(2)}
          padding={majorScale(2)}
        >
          <Pane flexDirection="column">
            <Heading marginBottom={majorScale(2)}>Settings</Heading>

            <Heading>{props.user.name}</Heading>
            
            <PlaidLink
              clientName="MOSS"
              env={props.plaidEnv}
              product={['transactions']}
              publicKey={props.plaidPublicKey}
              onExit={props.handleOnLinkExit}
              onSuccess={props.handleOnLinkSuccess}
              style={{
                outline: "none",
                background: "none",
                border: "none"
              }}
            >
              <Button appearance="primary">Add Account</Button>
            </PlaidLink>
          </Pane>
        </Card>
      </Pane>
    )
  }
}
