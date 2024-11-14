package br.com.josenaldo.wbu.services;

import br.com.josenaldo.wbu.entities.Account;
import br.com.josenaldo.wbu.entities.EntityId;
import br.com.josenaldo.wbu.exceptions.NotFoundException;
import br.com.josenaldo.wbu.repositories.AccountRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;
import java.util.Optional;


@Service
public class AccountService {

    private final AccountRepository accountRepository;

    public AccountService(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    public List<Account> findAll() {
        return accountRepository.findAll();
    }

    public Account findById(String accountId) {
        EntityId entityId = new EntityId(accountId);
        Optional<Account> account = accountRepository.findById(entityId);

        if(account.isEmpty()) {
            throw new NotFoundException("Account not found");
        }

        return account.get();
    }

    public void updateBalances(String accountIdFrom, String accountIdTo, BigDecimal balanceFrom, BigDecimal balanceTo) {
        EntityId entityIdFrom = new EntityId(accountIdFrom);
        EntityId entityIdTo = new EntityId(accountIdTo);

        Optional<Account> accountFrom = accountRepository.findById(entityIdFrom);
        Optional<Account> accountTo = accountRepository.findById(entityIdTo);

        if(accountFrom.isEmpty() || accountTo.isEmpty()) {
            throw new NotFoundException("Account not found");
        }

        Account accountFromEntity = accountFrom.get();
        Account accountToEntity = accountTo.get();

        accountFromEntity.setBalance(balanceFrom);
        accountToEntity.setBalance(balanceTo);

        accountRepository.save(accountFromEntity);
        accountRepository.save(accountToEntity);
    }
}
