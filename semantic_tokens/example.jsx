// 导入语句
import React, { useState, useEffect, useContext, useRef, useCallback, useMemo } from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';
import { BrowserRouter, Route, Switch, Link, useParams, useHistory } from 'react-router-dom';

// 上下文创建
const ThemeContext = React.createContext({
  theme: 'light',
  toggleTheme: () => {},
});

// 自定义Hook
function useLocalStorage(key, initialValue) {
  const [storedValue, setStoredValue] = useState(() => {
    try {
      const item = window.localStorage.getItem(key);
      return item ? JSON.parse(item) : initialValue;
    } catch (error) {
      console.error(error);
      return initialValue;
    }
  });

  const setValue = (value) => {
    try {
      const valueToStore = value instanceof Function ? value(storedValue) : value;
      setStoredValue(valueToStore);
      window.localStorage.setItem(key, JSON.stringify(valueToStore));
    } catch (error) {
      console.error(error);
    }
  };

  return [storedValue, setValue];
}

// 函数组件 - Button
function Button({ children, onClick, disabled = false, variant = 'primary', size = 'medium', icon }) {
  // 生成类名
  const buttonClass = `btn btn-${variant} btn-${size} ${disabled ? 'btn-disabled' : ''}`;

  return (
    <button
      className={buttonClass}
      onClick={onClick}
      disabled={disabled}
      aria-label={typeof children === 'string' ? children : undefined}
    >
      {icon && <span className='btn-icon'>{icon}</span>}
      <span className='btn-text'>{children}</span>
    </button>
  );
}

// Props 验证
Button.propTypes = {
  children: PropTypes.node.isRequired,
  onClick: PropTypes.func,
  disabled: PropTypes.bool,
  variant: PropTypes.oneOf(['primary', 'secondary', 'outlined', 'text']),
  size: PropTypes.oneOf(['small', 'medium', 'large']),
  icon: PropTypes.element,
};

// 主题提供者组件
function ThemeProvider({ children }) {
  const [theme, setTheme] = useState('light');

  const toggleTheme = useCallback(() => {
    setTheme((currentTheme) => (currentTheme === 'light' ? 'dark' : 'light'));
  }, []);

  const contextValue = useMemo(
    () => ({
      theme,
      toggleTheme,
    }),
    [theme, toggleTheme],
  );

  useEffect(() => {
    document.body.dataset.theme = theme;
  }, [theme]);

  return <ThemeContext.Provider value={contextValue}>{children}</ThemeContext.Provider>;
}

// 类组件示例
class UserCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isExpanded: false,
      isHovered: false,
    };

    this.toggleExpand = this.toggleExpand.bind(this);
    this.cardRef = React.createRef();
  }

  componentDidMount() {
    console.log('UserCard mounted');
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevState.isExpanded !== this.state.isExpanded) {
      console.log('Expansion state changed');
    }
  }

  componentWillUnmount() {
    console.log('UserCard will unmount');
  }

  toggleExpand() {
    this.setState((prevState) => ({
      isExpanded: !prevState.isExpanded,
    }));
  }

  handleMouseEnter = () => {
    this.setState({ isHovered: true });
  };

  handleMouseLeave = () => {
    this.setState({ isHovered: false });
  };

  render() {
    const { user, onDelete } = this.props;
    const { isExpanded, isHovered } = this.state;

    return (
      <div
        className={`user-card ${isExpanded ? 'expanded' : ''} ${isHovered ? 'hovered' : ''}`}
        ref={this.cardRef}
        onMouseEnter={this.handleMouseEnter}
        onMouseLeave={this.handleMouseLeave}
      >
        <div className='user-card-header'>
          <img src={user.avatar} alt={`${user.name}'s avatar`} className='avatar' />
          <h3>{user.name}</h3>
        </div>

        {isExpanded && (
          <div className='user-card-details'>
            <p>Email: {user.email}</p>
            <p>Phone: {user.phone}</p>
            <p>Position: {user.position}</p>
          </div>
        )}

        <div className='user-card-actions'>
          <Button onClick={this.toggleExpand} variant='outlined' size='small'>
            {isExpanded ? 'Show Less' : 'Show More'}
          </Button>

          <Button onClick={() => onDelete(user.id)} variant='text' size='small' icon={<TrashIcon />}>
            Delete
          </Button>
        </div>
      </div>
    );
  }
}

// 图标组件
const TrashIcon = () => (
  <svg viewBox='0 0 24 24' width='16' height='16' stroke='currentColor' fill='none'>
    <path d='M3 6h18M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2' />
    <line x1='10' y1='11' x2='10' y2='17' />
    <line x1='14' y1='11' x2='14' y2='17' />
  </svg>
);

// 用户列表组件（函数组件 + Hooks）
function UserList() {
  const [users, setUsers] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const { theme } = useContext(ThemeContext);
  const history = useHistory();

  // 获取用户数据
  useEffect(() => {
    async function fetchUsers() {
      setIsLoading(true);
      setError(null);

      try {
        const response = await fetch('https://api.example.com/users');
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        setUsers(data);
      } catch (error) {
        setError(error.message);
        console.error('Error fetching users:', error);
      } finally {
        setIsLoading(false);
      }
    }

    fetchUsers();

    // 清理函数
    return () => {
      console.log('UserList component unmounted');
    };
  }, []);

  // 处理删除用户
  const handleDeleteUser = useCallback((userId) => {
    setUsers((currentUsers) => currentUsers.filter((user) => user.id !== userId));
  }, []);

  // 处理用户点击
  const handleUserClick = useCallback(
    (userId) => {
      history.push(`/users/${userId}`);
    },
    [history],
  );

  // 渲染用户卡片
  const renderUserCards = () => {
    if (isLoading) {
      return <div className='loading'>Loading users...</div>;
    }

    if (error) {
      return <div className='error'>Error: {error}</div>;
    }

    if (users.length === 0) {
      return <div className='empty-state'>No users found</div>;
    }

    return (
      <div className={`user-grid theme-${theme}`}>
        {users.map((user) => (
          <UserCard key={user.id} user={user} onDelete={handleDeleteUser} onClick={() => handleUserClick(user.id)} />
        ))}
      </div>
    );
  };

  return (
    <div className='user-list-container'>
      <h2>User Directory</h2>
      {renderUserCards()}
    </div>
  );
}

// 用户详情页组件
function UserDetail() {
  const { userId } = useParams();
  const [user, setUser] = useState(null);
  const [isSaving, setIsSaving] = useState(false);
  const formRef = useRef(null);

  useEffect(() => {
    async function fetchUserDetail() {
      try {
        const response = await fetch(`https://api.example.com/users/${userId}`);
        const userData = await response.json();
        setUser(userData);
      } catch (error) {
        console.error('Error fetching user details:', error);
      }
    }

    if (userId) {
      fetchUserDetail();
    }
  }, [userId]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsSaving(true);

    try {
      // 获取表单数据
      const formData = new FormData(formRef.current);
      const userData = Object.fromEntries(formData);

      // 发送更新请求
      await fetch(`https://api.example.com/users/${userId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
      });

      // 更新本地状态
      setUser((prevUser) => ({ ...prevUser, ...userData }));
      alert('User updated successfully!');
    } catch (error) {
      console.error('Error updating user:', error);
      alert('Failed to update user');
    } finally {
      setIsSaving(false);
    }
  };

  if (!user) {
    return <div>Loading user details...</div>;
  }

  return (
    <div className='user-detail'>
      <h2>Edit User: {user.name}</h2>

      <form ref={formRef} onSubmit={handleSubmit} className='user-form'>
        <div className='form-group'>
          <label htmlFor='name'>Name</label>
          <input type='text' id='name' name='name' defaultValue={user.name} required />
        </div>

        <div className='form-group'>
          <label htmlFor='email'>Email</label>
          <input type='email' id='email' name='email' defaultValue={user.email} required />
        </div>

        <div className='form-group'>
          <label htmlFor='phone'>Phone</label>
          <input type='tel' id='phone' name='phone' defaultValue={user.phone} />
        </div>

        <div className='form-actions'>
          <Button type='submit' disabled={isSaving} variant='primary'>
            {isSaving ? 'Saving...' : 'Save Changes'}
          </Button>

          <Link to='/users'>
            <Button variant='outlined'>Cancel</Button>
          </Link>
        </div>
      </form>
    </div>
  );
}

// 主应用组件
function App() {
  const [preferences, setPreferences] = useLocalStorage('app-preferences', {
    notifications: true,
    compactView: false,
  });

  return (
    <BrowserRouter>
      <ThemeProvider>
        <div className='app'>
          <header className='app-header'>
            <h1>React User Management</h1>
            <nav>
              <ul className='nav-links'>
                <li>
                  <Link to='/'>Home</Link>
                </li>
                <li>
                  <Link to='/users'>Users</Link>
                </li>
                <li>
                  <Link to='/settings'>Settings</Link>
                </li>
              </ul>
            </nav>
            <ThemeToggle />
          </header>

          <main className='app-content'>
            <Switch>
              <Route exact path='/'>
                <Home />
              </Route>
              <Route exact path='/users'>
                <UserList />
              </Route>
              <Route path='/users/:userId'>
                <UserDetail />
              </Route>
              <Route path='/settings'>
                <Settings preferences={preferences} setPreferences={setPreferences} />
              </Route>
              <Route path='*'>
                <NotFound />
              </Route>
            </Switch>
          </main>

          <footer className='app-footer'>
            <p>&copy; 2025 User Management System</p>
          </footer>
        </div>
      </ThemeProvider>
    </BrowserRouter>
  );
}

// 主题切换组件
function ThemeToggle() {
  const { theme, toggleTheme } = useContext(ThemeContext);

  return (
    <button
      className='theme-toggle'
      onClick={toggleTheme}
      aria-label={`Switch to ${theme === 'light' ? 'dark' : 'light'} mode`}
    >
      {theme === 'light' ? '🌙' : '☀️'}
    </button>
  );
}

// 其他页面组件
function Home() {
  return (
    <div className='home-page'>
      <h2>Welcome to User Management</h2>
      <p>This application allows you to manage users in your system.</p>
      <Link to='/users'>
        <Button variant='primary' size='large'>
          View Users
        </Button>
      </Link>
    </div>
  );
}
