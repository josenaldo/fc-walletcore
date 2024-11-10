package br.com.josenaldo.wbu.controllers;

import br.com.josenaldo.wbu.entities.Account;
import br.com.josenaldo.wbu.exceptions.InvalidIdException;
import br.com.josenaldo.wbu.exceptions.NotFoundException;
import br.com.josenaldo.wbu.services.AccountService;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping(path = "/balances", produces = MediaType.APPLICATION_JSON_VALUE)
public class AccountsController {

    private final AccountService accountService;

    public AccountsController(AccountService accountService) {
        this.accountService = accountService;
    }

    @GetMapping(value={"", "/"})
    public ResponseEntity<List<Account>> getBalances() {
        List<Account> accounts = this.accountService.findAll();
        return ResponseEntity.ok(accounts) ;
    }

    @GetMapping("/{account_id}")
    public ResponseEntity<?> getBalance(@PathVariable String account_id) {
        try {
            Account account = this.accountService.findById(account_id);
            return ResponseEntity.ok(account);
        }catch (NotFoundException e) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(e.getMessage());
        }catch(InvalidIdException e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(e.getMessage());
        }
    }
}
