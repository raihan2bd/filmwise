const Button = ({btnClass, type = "button", onClick, children }) => {
  const btnClasses = btnClass ? `btn-default ${btnClass}` : `btn-default`;
  return (
    <button className={btnClasses} type={type} onClick={onClick}>
      {children}
    </button>
  );
};

export default Button;
