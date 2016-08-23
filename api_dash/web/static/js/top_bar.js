import React from 'react'
import AppBar from 'material-ui/AppBar'
import Drawer from 'material-ui/Drawer'
import baseTheme from 'material-ui/styles/baseThemes/lightBaseTheme'
import getMuiTheme from 'material-ui/styles/getMuiTheme'
import MenuItem from 'material-ui/MenuItem'

export default class TopBar extends React.Component {

  constructor(props) {
    super(props)
    this.state = {openLeftBar: false}
  }

  getChildContext() {
    return { muiTheme: getMuiTheme(baseTheme) }
  }

  toggleLeftMenu() {
    this.setState({openLeftBar: ! this.state.openLeftBar})
  }

  linkTo() {
    debugger
  }

  render() {
    return <div>
      <AppBar
        title='Api Dash'
        onLeftIconButtonTouchTap={this.toggleLeftMenu.bind(this)}></AppBar>
      <Drawer open={this.state.openLeftBar}>
        <AppBar onLeftIconButtonTouchTap={this.toggleLeftMenu.bind(this)}
          title='Api Dash'>
        </AppBar>
        <MenuItem href='/users'>Users</MenuItem>
        <MenuItem onTouchTap={this.linkTo}>Apps</MenuItem>
      </Drawer>
    </div>
  }
}

TopBar.childContextTypes = {
   muiTheme: React.PropTypes.object.isRequired,
}
