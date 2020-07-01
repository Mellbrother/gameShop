import axios from 'axios'   // apiとの通信を簡単にする
import React, { Component } from 'react';
import web3 from './Web3'   // ethereumとの通信用
import GameShop from './contract/GameShop'
import styled, { createGlobalStyle } from 'styled-components'   // CSS in JS
import reset from 'styled-reset'
import { ArrowLeft, Plus } from 'styled-icons/fa-solid'

// GameShopのアドレスに変更する必要あり
const address = '0xE39BcE1506dCE33D482F86051B3408d09F2cB06d'
const gameShop = new web3.eth.Contract(GameShop.abi, address)

class App extends Component {
  constructor (props) {
    super(props)
    this.state = {
      price: 0,
      items: [],
      isShowModal: false    // 商品登録で使う    
    }
  }

  async componentDidMount() {
    // Itemi一覧APIのURL
    const res = await axios.get('http://localhost:8080/items')
    this.setState({ items: res.data })
  }

  handleKeyUp = e => {
    this.setState({ price: e.target.value })
  }

  handleRegisterItem = async () => {
    const [ethaddress] = await web3.eth.getAccounts()

    try {
      gameShop.methods.registerItem(200).send({
        from: ethaddress,
      })
    } catch (error) {
      console.log(error)
    }
  }

  handleCancel = () => {
    this.setState({ price: '', isShowModal: false })
  }

  handleBack = () => {
    this.setState({ price: '', isShowModal: false })
  }

  handleModal = () => {
    this.setState({ isShowModal: true })
  }


  render() {
    const { items } = this.state


    // if文に直した方が可読性が上がる
    return (
      <>
        <GlobalStyle />
        {
          !this.state.isShowModal ? (
            // Items
            <>
              <Header>
                <h1>Items</h1>
              </Header>
              <Subtitle>
                <p>商品一覧</p>
                <ModalButton onClick={this.handleModal}>
                  <StyledPlus />
                  Register Item
                </ModalButton>
              </Subtitle>
              {
                // console.log("aiueo") &&
                items && items.length !== 0 &&
                  items.map(item => (
                    <Item>
                      <ItemName>{item.name}</ItemName>
                      <ItemPrice>
                        <Label>価格</Label>
                        <ItemPrice>{item.price}</ItemPrice>
                      </ItemPrice>
                    </Item>  
                  ))
              }
            </>
          ) : (
            // Create Room
            <>
              <Header>
                <BackButton onClick={this.handleBack}>
                  <ArrowLeft />
                </BackButton>
                <h1>Register Item</h1>
              </Header>
              <Contents>
                <p>価格を入力</p>
                <input onKeyUp={this.handleKeyUp} placeholder='Item Price (YEN)' />
              </Contents>
              <Footer>
                <button onClick={this.handleCancel} className='cancel'>キャンセル</button>
                <button onClick={this.handleRegisterItem} className='save'>保存</button>
              </Footer>
            </>
          )
        }
      </>
    )
  }
}

export default App;

const GlobalStyle = createGlobalStyle`
  ${reset}
  body {
    color: #4a4a4a;
  }
`

const Header = styled.header`
  background: linear-gradient(to right, #9C56D3, #804AAA);
  padding: 20px;

  h1 {
    color: #fff;
    text-align: center;
  }
`

const BackButton = styled.div`
  position: absolute;
  width: 18px;
  color: #fff;
  cursor: pointer;
`

const Contents = styled.div`
  padding: 40px;

  > p {
    margin-bottom: 10px;
  }

  > input {
    align-items: center;
    border-radius: 4px;
    box-shadow: none;
    font-size: 1rem;
    height: 2.25em;
    line-height: 1.5;
    padding: 5px 10px;
    position: relative;
    vertical-align: top;
    border: 1px solid #ccc;
    display: block;
    width: 100%;
    box-sizing: border-box;
  }
`

const Subtitle = styled.div`
  background: #f8f8f8;
  display: flex;
  padding: 10px;
  justify-content: space-between;

  > p {
    line-height: 3;
    font-size: 12px;
  }
`

const ModalButton = styled.button`
  background-color: #23d160;
  border-color: transparent;
  cursor: pointer;
  color: #fff;
  padding: 10px 15px;
  border-radius: 50px;
  font-size: 16px;
  font-weight: 100;
`

const StyledPlus = styled(Plus)`
  width: 16px;
  margin-right: 10px;
`

const Item = styled.div`
  border-bottom: 1px solid #ccc;
  padding: 20px;
`

const ItemName = styled.div`
  color: #363636;
  font-size: 20px;
  margin-bottom: 10px;
`

const ItemPrice = styled.div`
  display: flex;
`

const Label = styled.div`
  font-size: 11px;
  padding: 5px 7px;
  margin-right: 8px;
  border-radius: 5px;
  color: #4A4A4A;
  background: #eee;
  white-space: nowrap;
`

const Footer = styled.footer`
  border-top: 1px solid #ccc;
  background: #fafafa;
  display: -ms-flexbox;
  display: flex;
  padding: 10px;
  justify-content: center;

  > button {
    color: #fff;
    font-weight: bold;
    min-width: 120px;
    padding: 8px 24px;
    height: auto;
    font-size: 15px;
    border: none;
    border-radius: 4px;
    
    &.cancel {
      background: #F5A623;
      margin-right: 15px;
    }

    &.save {
      background: #6FBE4E;
    }
  }
`