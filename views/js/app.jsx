class App extends React.Component {
    render() {
      if (this.loggedIn) {
        return (<LoggedIn />);
      } else {
        return (<Home />);
      }
    }
  }

  class Home extends React.Component {
    render() {
      return (
        <div className="container">
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>Go Cast </h1>
            <p> For all your Safe File Transfer needs </p>
            <p>Sign in to begin </p>
            <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
          </div>
        </div>
      )
    }
  }

  ReactDOM.render(<App />, document.getElementById('app'));