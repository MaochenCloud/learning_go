//dao:
return errors.Wrapf(code.noFound, fmt.Sprintf("sql: %s error: %v", sql, err))

//up

if errors.Is(err, code.noFound) {
}

