import React from 'react';
import Article from '../article/Article';
import axios from 'axios';

class ArticleList extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      articles: [
        {title: "title-1", body: "body-1" },
        {title: "title-2", body: "body-2" },
        {title: "title-3", body: "body-3" },
      ]
    };
    this.handleGetArticles = this.handleGetArticles.bind(this);
  }

  handleGetArticles() {
    axios.get("http://localhost:1991/articles").then((result) => {
      const data = result.data;

      this.setState({articles: data});
    }).catch(() => {
      console.log("tsusin miss");
    })
  }

  render() {
    return (
      <div>
        <h1>記事リスト</h1>
        <div>
          <ul>
            {
              this.state.articles.map((article, i) => {
                return <Article key={i} title={article.title} body={article.body} />
              })
            }
          </ul>
        </div>
        <button onClick={() => this.handleGetArticles()}>fetch!!</button>
      </div>
    );
  }
}

export default ArticleList
