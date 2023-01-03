import { css, cx } from '@emotion/css';
import styled from '@emotion/styled';
import React from 'react';
import { HiOutlineDatabase } from 'react-icons/hi';
import { HiOutlineWrenchScrewdriver } from 'react-icons/hi2';
import { Link } from 'react-router-dom';

const Nav = styled.nav(`
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 20px;
  font-size: 16px;
  color: #16161a;
  flex-grow: 1;
  justify-content: flex-start;
  overflow-x: hidden;
  margin-top: 34px;
  margin-bottom: 30px;
`);

const Aside = styled.aside(`
  position: relative;
  z-index: 1;
  display: flex;
  flex-grow: 0;
  flex-shrink: 0;
  flex-direction: column;
  align-items: center;
  width: 75px;
  padding: 20px 0 20px;
  background: #ffffff;
  border-right: 1px solid #dee3e8;
`);

const Logo = styled(Link)`
  position: relative;
  vertical-align: top;
  display: inline-flex;
  align-items: center;
  gap: 10px;
  font-size: 20px;
  text-decoration: none;
  color: inherit;
  user-select: none;
`;

const SidebarLink = styled(Link)`
  position: relative;
  outline: 0;
  cursor: pointer;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  text-align: center;
  justify-content: center;
  user-select: none;
  color: inherit;
  min-width: 45px;
  min-height: 45px;
  border: 2px solid transparent;
  border-radius: 12px;
  transition: background 0.3s, border 0.3s;

  :hover {
    background: #ebeff2;
  }
`;

const sidebarActiveCss = css`
  background: #ffffff;
  border-color: #16161a !important;
`;

const layoutCss = css`
  display: flex;
  width: 100%;
  height: 100vh;
`;

const navIconCss = css`
  font-size: 24px;
  line-height: 1;
`;

const Figure = css``;

export const LayoutComponent = ({ children }) => {
  return (
    <div className={layoutCss}>
      <Aside>
        <Logo to="/">WA</Logo>
        <Nav>
          <SidebarLink to="/" className={cx({ [sidebarActiveCss]: true })}>
            <HiOutlineDatabase className={navIconCss} />
          </SidebarLink>
          <SidebarLink to="/" className={cx({ [sidebarActiveCss]: false })}>
            <HiOutlineWrenchScrewdriver className={navIconCss} />
          </SidebarLink>
        </Nav>
        <Figure>Profile</Figure>
      </Aside>
      {children}
    </div>
  );
};
