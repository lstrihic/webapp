import React from 'react';
import { css } from '@emotion/react';
import { Helmet } from 'react-helmet-async';

export const HomePage = () => {
  return (
    <>
      <Helmet title="Home Page" />
      <div className="container">
        <div className="columns">
          <div className="column">
            First column
            <div
              css={css`
                background-color: hotpink;
                &:hover {
                  background-color: red;
                  cursor: pointer;
                  color: white;
                }
              `}
            >
              Styled
            </div>
          </div>
          <div className="column">Second column</div>
          <div className="column">Third column</div>
          <div className="column">Fourth column</div>
        </div>
      </div>
    </>
  );
};
