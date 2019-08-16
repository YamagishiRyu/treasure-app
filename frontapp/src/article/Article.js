import React from 'react';

class Article extends React.Component {
  render() {
    return (
      <div>
        <div>
          {this.props.title}
        </div>
        <div>
          {this.props.body}
        </div>
      </div>
    )
  }
}

export default Article
