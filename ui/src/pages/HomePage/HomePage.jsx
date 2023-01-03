import { LayoutComponent } from '@components';
import { css } from '@emotion/react';
import styled from '@emotion/styled';
import React from 'react';
import { Helmet } from 'react-helmet-async';

const Main = styled.main(`
  position: relative;
  display: block;
  width: 100%;
  flex-grow: 1;
  padding: 25px 30px;
`);

const Header = styled.header(`
  display: flex;
  align-items: center;
  width: 100%;
  min-height: 40px;
  gap: 15px;
  margin: 0 0 30px;
`);

const HeaderNav = styled.nav(`
  display: flex;
  align-items: center;
  gap: 30px;
  color: #adb3b8;
`);

const HeaderNavItem = styled.div(`
  position: relative;
  margin: 0;
  line-height: 1;
  font-weight: 400;
  font-size: 20px;
  
  :after {
    content: "/";
    position: absolute;
    right: -20px;
    top: 0;
    width: 10px;
    text-align: center;
    pointer-events: none;
    opacity: .4;
  }
  :last-child {
    :after { display: none; }
    word-break: break-word;
    color: #16161a;
  }
`);

export const HomePage = () => {
  return (
    <LayoutComponent>
      <Helmet title="Home Page" />
      <Main>
        <Header>
          <HeaderNav>
            <HeaderNavItem>Dashboard</HeaderNavItem>
            <HeaderNavItem>home</HeaderNavItem>
          </HeaderNav>
        </Header>
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
      </Main>
    </LayoutComponent>
  );
};
